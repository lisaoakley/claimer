package translate_test

import (
	. "github.com/mdelillo/claimer/translate"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
)

var _ = Describe("Translate", func() {
	Describe("T", func() {
		var (
			language1TranslationPath string
			language2TranslationPath string
		)

		BeforeEach(func() {
			language1Translation, err := ioutil.TempFile("", "claimer-translate")
			Expect(err).NotTo(HaveOccurred())
			language1TranslationPath = language1Translation.Name()
			Expect(ioutil.WriteFile(language1TranslationPath, []byte("some: {field: field-in-language1}"), 0644)).To(Succeed())

			language2Translation, err := ioutil.TempFile("", "claimer-translate")
			Expect(err).NotTo(HaveOccurred())
			language2TranslationPath = language2Translation.Name()
			Expect(ioutil.WriteFile(language2TranslationPath, []byte("some: {field: field-in-language2}"), 0644)).To(Succeed())
		})

		AfterEach(func() {
			os.RemoveAll(language1TranslationPath)
			os.RemoveAll(language2TranslationPath)
		})

		It("translates based on the loaded file", func() {
			Expect(LoadTranslationFile(language1TranslationPath)).To(Succeed())
			Expect(T("some.field")).To(Equal("field-in-language1"))
			Expect(LoadTranslationFile(language2TranslationPath)).To(Succeed())
			Expect(T("some.field")).To(Equal("field-in-language2"))
		})
	})
})
