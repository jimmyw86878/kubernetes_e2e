/*
This is an example from Ginkgo official website. Just a demonstration for Ginkgo.
*/

package k8s_test

import (
	"github.com/onsi/ginkgo"
)

type Book struct {
	Title  string
	Author string
	Pages  int64
}

var _ = ginkgo.Describe("Book", func() {
	// var (
	// 	longBook  Book
	// 	shortBook Book
	// )

	// ginkgo.BeforeEach(func() {
	// 	longBook = Book{
	// 		Title:  "Les Miserables",
	// 		Author: "Victor Hugo",
	// 		Pages:  2783,
	// 	}

	// 	shortBook = Book{
	// 		Title:  "Fox In Socks",
	// 		Author: "Dr. Seuss",
	// 		Pages:  24,
	// 	}
	// })

	// ginkgo.Describe("Categorizing book length", func() {
	// 	ginkgo.Context("With more than 300 pages", func() {
	// 		ginkgo.It("should be a novel", func() {
	// 			gomega.Expect(longBook.CategoryByLength()).To(gomega.Equal("NOVEL"))
	// 		})
	// 	})

	// 	ginkgo.Context("With fewer than 300 pages", func() {
	// 		ginkgo.It("should be a short story", func() {
	// 			gomega.Expect(shortBook.CategoryByLength()).To(gomega.Equal("SHORT STORY"))
	// 		})
	// 	})
	// })
})

func (b *Book) CategoryByLength() string {
	if b.Pages > 300 {
		return "NOVEL"
	}
	return "SHORT STORY"
}
