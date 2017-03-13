package locker_test

import (
	. "github.com/mdelillo/claimer/locker"

	"errors"
	"github.com/mdelillo/claimer/locker/lockerfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"path/filepath"
)

var _ = Describe("Locker", func() {
	Describe("ClaimLock", func() {
		var (
			fs      *lockerfakes.FakeFs
			gitRepo *lockerfakes.FakeGitRepo
		)

		BeforeEach(func() {
			fs = new(lockerfakes.FakeFs)
			gitRepo = new(lockerfakes.FakeGitRepo)
		})

		It("claims the lock file in the git repo", func() {
			pool := "some-pool"
			gitDir := "some-dir"
			lock := "some-lock"

			gitRepo.DirReturns(gitDir)
			fs.LsReturns([]string{lock}, nil)

			locker := NewLocker(fs, gitRepo)
			Expect(locker.ClaimLock(pool)).To(Succeed())

			Expect(gitRepo.CloneOrPullCallCount()).To(Equal(1))

			Expect(fs.LsCallCount()).To(Equal(1))
			Expect(fs.LsArgsForCall(0)).To(Equal(filepath.Join(gitDir, pool, "unclaimed")))

			Expect(fs.MvCallCount()).To(Equal(1))
			oldPath, newPath := fs.MvArgsForCall(0)
			Expect(oldPath).To(Equal(filepath.Join(gitDir, pool, "unclaimed", lock)))
			Expect(newPath).To(Equal(filepath.Join(gitDir, pool, "claimed", lock)))

			Expect(gitRepo.CommitAndPushCallCount()).To(Equal(1))
			Expect(gitRepo.CommitAndPushArgsForCall(0)).To(Equal("Claimer claiming " + pool))
		})

		Context("when cloning the repo fails", func() {
			It("returns an error", func() {
				gitRepo.CloneOrPullReturns(errors.New("some-error"))

				locker := NewLocker(fs, gitRepo)
				Expect(locker.ClaimLock("")).To(MatchError("some-error"))
			})
		})

		Context("when listing files fails", func() {
			It("returns an error", func() {
				fs.LsReturns(nil, errors.New("some-error"))

				locker := NewLocker(fs, gitRepo)
				Expect(locker.ClaimLock("")).To(MatchError("some-error"))
			})
		})

		Context("when there are no unclaimed locks", func() {
			It("returns an error", func() {
				pool := "some-pool"

				fs.LsReturns([]string{}, nil)

				locker := NewLocker(fs, gitRepo)
				Expect(locker.ClaimLock(pool)).To(MatchError("no unclaimed locks for pool " + pool))
			})
		})

		Context("when there are multiple unclaimed locks", func() {
			It("returns an error", func() {
				pool := "some-pool"

				fs.LsReturns([]string{"some-lock", "some-other-lock"}, nil)

				locker := NewLocker(fs, gitRepo)
				Expect(locker.ClaimLock(pool)).To(MatchError("too many unclaimed locks for pool " + pool))
			})
		})

		Context("when moving the file fails", func() {
			It("returns an error", func() {
				fs.LsReturns([]string{"some-lock"}, nil)
				fs.MvReturns(errors.New("some-error"))

				locker := NewLocker(fs, gitRepo)
				Expect(locker.ClaimLock("")).To(MatchError("some-error"))
			})
		})

		Context("when pushing fails", func() {
			It("returns an error", func() {
				fs.LsReturns([]string{"some-lock"}, nil)
				gitRepo.CommitAndPushReturns(errors.New("some-error"))

				locker := NewLocker(fs, gitRepo)
				Expect(locker.ClaimLock("")).To(MatchError("some-error"))
			})
		})
	})
})