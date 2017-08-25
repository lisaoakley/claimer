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
		var translationsPath string

		AfterEach(func() {
			os.RemoveAll(translationsPath)
		})

		Context("multiple translation files", func() {
			var otherTranslationsPath string

			BeforeEach(func() {
				translationsPath = writeTranslationFile("key: field-in-language1\nkey1: field1")
				otherTranslationsPath = writeTranslationFile("key: field-in-language2\nkey2: field2")
				Expect(LoadTranslationFile(translationsPath)).To(Succeed())
				Expect(LoadTranslationFile(otherTranslationsPath)).To(Succeed())
			})

			AfterEach(func() {
				os.RemoveAll(otherTranslationsPath)
			})

			It("translates with the last loaded file taking precedence", func() {
				Expect(T("key", nil)).To(Equal("field-in-language2"))
				Expect(T("key1", nil)).To(Equal("field1"))
				Expect(T("key2", nil)).To(Equal("field2"))
			})
		})

		Context("nested keys", func() {
			BeforeEach(func() {
				translationsPath = writeTranslationFile("some: {nested: {key: some-value}}")
				Expect(LoadTranslationFile(translationsPath)).To(Succeed())
			})

			It("translates using the nested value", func() {
				Expect(T("some.nested.key", nil)).To(Equal("some-value"))
			})
		})

		Context("passing variables to translation", func() {
			BeforeEach(func() {
				translationsPath = writeTranslationFile("key: some {{.var1}} interpolated {{.var2}} string")
				Expect(LoadTranslationFile(translationsPath)).To(Succeed())
			})

			It("interpolates the variables into the translated string", func() {
				vars := map[string]string{
					"var1": "value1",
					"var2": "value2",
				}
				Expect(T("key", vars)).To(Equal("some value1 interpolated value2 string"))
			})

			It("does not interpolate variables that are not passed in", func() {
				vars := map[string]string{
					"var2": "value2",
				}
				Expect(T("key", vars)).To(Equal("some {{.var1}} interpolated value2 string"))
			})
		})

		Context("translation file errors", func() {
			Context("when the translation file does not exist", func() {
				It("returns an error", func() {
					Expect(LoadTranslationFile("some-bad-path")).To(MatchError("failed to read file: some-bad-path"))
				})
			})

			Context("when the translation file is not valid YAML", func() {
				BeforeEach(func() {
					translationsPath = writeTranslationFile("some-invalid-yaml")
				})

				It("returns an error", func() {
					Expect(LoadTranslationFile(translationsPath)).To(MatchError("failed to parse YAML: some-invalid-yaml"))
				})
			})
		})

		Context("key errors", func() {
			Context("when the key does not exit", func() {
				BeforeEach(func() {
					translationsPath = writeTranslationFile("---")
				})

				It("returns the untranslated key", func() {
					Expect(LoadTranslationFile(translationsPath)).To(Succeed())
					Expect(T("missingkey", nil)).To(Equal("missingkey"))
				})
			})

			Context("when the nested key does not exit", func() {
				BeforeEach(func() {
					translationsPath = writeTranslationFile("nested: {key: some-value}")
				})

				It("returns the untranslated key", func() {
					Expect(LoadTranslationFile(translationsPath)).To(Succeed())
					Expect(T("nested.missingkey", nil)).To(Equal("nested.missingkey"))
				})
			})

			Context("when a value is a string instead of a nested map", func() {
				BeforeEach(func() {
					translationsPath = writeTranslationFile("nested: {key: value}")
				})

				It("returns the untranslated key", func() {
					Expect(LoadTranslationFile(translationsPath)).To(Succeed())
					Expect(T("nested.key.otherkey", nil)).To(Equal("nested.key.otherkey"))
				})
			})

			Context("when a value is not a string or map", func() {
				BeforeEach(func() {
					translationsPath = writeTranslationFile("key: [not-a-string-or-map]")
				})

				It("returns the untranslated key", func() {
					Expect(LoadTranslationFile(translationsPath)).To(Succeed())
					Expect(T("key", nil)).To(Equal("key"))
				})
			})
		})
	})
})

func writeTranslationFile(translations string) string {
	file, err := ioutil.TempFile("", "claimer-translate")
	Expect(err).NotTo(HaveOccurred())
	Expect(ioutil.WriteFile(file.Name(), []byte(translations), 0644)).To(Succeed())
	return file.Name()
}