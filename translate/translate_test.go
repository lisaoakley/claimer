package translate_test

import (
	. "github.com/mdelillo/claimer/translate"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
)

var _ = Describe("Translate", func() {
	Describe("Translation", func() {
		var (
			language1Translations    string
			language2Translations    string
			language1TranslationPath string
			language2TranslationPath string
		)

		JustBeforeEach(func() {
			file1, err := ioutil.TempFile("", "claimer-translate")
			Expect(err).NotTo(HaveOccurred())
			language1TranslationPath = file1.Name()
			Expect(ioutil.WriteFile(language1TranslationPath, []byte(language1Translations), 0644)).To(Succeed())

			file2, err := ioutil.TempFile("", "claimer-translate")
			Expect(err).NotTo(HaveOccurred())
			language2TranslationPath = file2.Name()
			Expect(ioutil.WriteFile(language2TranslationPath, []byte(language2Translations), 0644)).To(Succeed())
		})

		AfterEach(func() {
			os.RemoveAll(language1TranslationPath)
			os.RemoveAll(language2TranslationPath)
		})

		Context("multiple translation files", func() {
			BeforeEach(func() {
				language1Translations = "key: field-in-language1"
				language2Translations = "key: field-in-language2"
			})

			It("translates based on the loaded file", func() {
				Expect(LoadTranslationFile(language1TranslationPath)).To(Succeed())
				Expect(T("key")).To(Equal("field-in-language1"))
				Expect(LoadTranslationFile(language2TranslationPath)).To(Succeed())
				Expect(T("key")).To(Equal("field-in-language2"))
			})
		})

		Context("nested keys", func() {
			BeforeEach(func() {
				language1Translations = "some: {nested: {key: some-value}}"
			})

			It("translates using the nested value", func() {
				Expect(LoadTranslationFile(language1TranslationPath)).To(Succeed())
				Expect(T("some.nested.key")).To(Equal("some-value"))
			})
		})

		Context("when the translation file does not exist", func() {
			It("returns an error", func() {
				Expect(LoadTranslationFile("some-bad-path")).To(MatchError("failed to read file: some-bad-path"))
			})
		})

		Context("when the translation file is not valid YAML", func() {
			BeforeEach(func() {
				language1Translations = "some-invalid-yaml"
			})

			It("returns an error", func() {
				Expect(LoadTranslationFile(language1TranslationPath)).To(MatchError("failed to parse YAML: some-invalid-yaml"))
			})
		})

		Context("when the key does not exit", func() {
			BeforeEach(func() {
				language1Translations = "---"
			})

			It("returns an error", func() {
				Expect(LoadTranslationFile(language1TranslationPath)).To(Succeed())
				_, err := T("missingkey")
				Expect(err).To(MatchError("could not find key: missingkey"))
			})
		})

		Context("when the nested key does not exit", func() {
			BeforeEach(func() {
				language1Translations = "nested: {key: some-value}"
			})

			It("returns an error", func() {
				Expect(LoadTranslationFile(language1TranslationPath)).To(Succeed())
				_, err := T("nested.missingkey")
				Expect(err).To(MatchError("could not find key: nested.missingkey"))
			})
		})

		Context("when a value is a string instead of a nested map", func() {
			BeforeEach(func() {
				language1Translations = "nested: {key: value}"
			})

			It("returns an error", func() {
				Expect(LoadTranslationFile(language1TranslationPath)).To(Succeed())
				_, err := T("nested.key.otherkey")
				Expect(err).To(MatchError("found string instead of map for key: nested.key"))
			})
		})

		Context("when a value is not a string or map", func() {
			BeforeEach(func() {
				language1Translations = "key: [not-a-string-or-map]"
			})

			It("returns an error", func() {
				Expect(LoadTranslationFile(language1TranslationPath)).To(Succeed())
				_, err := T("key")
				Expect(err).To(MatchError("could not convert value to map for key: key"))
			})
		})
	})
})
