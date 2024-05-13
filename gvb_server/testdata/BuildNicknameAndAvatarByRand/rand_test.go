package randomname

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"path"
	"testing"
	"unicode/utf8"

	"github.com/DanPlayer/randomname"
	"github.com/disintegration/letteravatar"
	"github.com/golang/freetype"
)

func TestRandomNameAndAvatar(t *testing.T) {
	name := randomname.GenerateName()
	fmt.Println("Name:", name)
	names := []rune(name)

	DrawImage(string(names[0]), "D:\\gin-vue-blog\\gvb_server\\uploads\\chat_avatar")
}

func DrawImage(name string, dir string) {
	fontFile, _ := os.ReadFile("D:\\gin-vue-blog\\gvb_server\\uploads\\system\\方正综艺简体.TTF")
	font, err := freetype.ParseFont(fontFile)
	if err != nil {
		fmt.Println(err)
		return

	}
	options := &letteravatar.Options{
		Font: font,
	}
	firstLetter, _ := utf8.DecodeRuneInString(name)
	img, err := letteravatar.Draw(140, firstLetter, options)
	if err != nil {
		log.Fatal(err)
	}
	filePath := path.Join(dir, name+".png")
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}

}

func TestGenerateNameAvater(t *testing.T) {
	dir := "D:\\gin-vue-blog\\gvb_server\\uploads\\chat_avatar"
	for _, s := range randomname.AdjectiveSlice {
		DrawImage(string([]rune(s)[0]), dir)
	}
	for _, s := range randomname.PersonSlice {
		DrawImage(string([]rune(s)[0]), dir)
	}

}
