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
				language1Translations = "key: field-in-language1\nkey1: field1"
				language2Translations = "key: field-in-language2\nkey2: field2"
			})

			It("translates with the last loaded file taking precedence", func() {
				Expect(LoadTranslationFile(language1TranslationPath)).To(Succeed())
				Expect(LoadTranslationFile(language2TranslationPath)).To(Succeed())
				Expect(T("key", nil)).To(Equal("field-in-language2"))
				Expect(T("key1", nil)).To(Equal("field1"))
				Expect(T("key2", nil)).To(Equal("field2"))
			})
		})

		Context("nested keys", func() {
			BeforeEach(func() {
				language1Translations = "some: {nested: {key: some-value}}"
			})

			It("translates using the nested value", func() {
				Expect(LoadTranslationFile(language1TranslationPath)).To(Succeed())
				Expect(T("some.nested.key", nil)).To(Equal("some-value"))
			})
		})

		Context("passing variables to translation", func() {
			BeforeEach(func() {
				language1Translations = "key: some {{.var1}} interpolated {{.var2}} string"
			})

			It("interpolates the variables into the translated string", func() {
				Expect(LoadTranslationFile(language1TranslationPath)).To(Succeed())
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
					language1Translations = "some-invalid-yaml"
				})

				It("returns an error", func() {
					Expect(LoadTranslationFile(language1TranslationPath)).To(MatchError("failed to parse YAML: some-invalid-yaml"))
				})
			})
		})

		Context("key errors", func() {
			Context("when the key does not exit", func() {
				BeforeEach(func() {
					language1Translations = "---"
				})

				It("returns the untranslated key", func() {
					Expect(LoadTranslationFile(language1TranslationPath)).To(Succeed())
					Expect(T("missingkey", nil)).To(Equal("missingkey"))
				})
			})

			Context("when the nested key does not exit", func() {
				BeforeEach(func() {
					language1Translations = "nested: {key: some-value}"
				})

				It("returns the untranslated key", func() {
					Expect(LoadTranslationFile(language1TranslationPath)).To(Succeed())
					Expect(T("nested.missingkey", nil)).To(Equal("nested.missingkey"))
				})
			})

			Context("when a value is a string instead of a nested map", func() {
				BeforeEach(func() {
					language1Translations = "nested: {key: value}"
				})

				It("returns the untranslated key", func() {
					Expect(LoadTranslationFile(language1TranslationPath)).To(Succeed())
					Expect(T("nested.key.otherkey", nil)).To(Equal("nested.key.otherkey"))
				})
			})

			Context("when a value is not a string or map", func() {
				BeforeEach(func() {
					language1Translations = "key: [not-a-string-or-map]"
				})

				It("returns the untranslated key", func() {
					Expect(LoadTranslationFile(language1TranslationPath)).To(Succeed())
					Expect(T("key", nil)).To(Equal("key"))
				})
			})
		})
	})
})
