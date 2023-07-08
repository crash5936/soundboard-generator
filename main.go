package main

import (
	"html/template"
	"io/fs"
	"os"
	"regexp"
)

type Sound struct {
	Name, Description string
}

func main() {
	sounds, err := os.ReadDir("./sounds/")

	check_error(err)

	sound_names := get_sound_names(sounds)

	var sound_structs []Sound

	for _, name := range sound_names {
		sound_structs = append(sound_structs, Sound{name, name})
	}

	check_error(err)

	tmpl_str, err := os.ReadFile("templates/test.tmpl")

	tmpl, err := template.New("soundboard").Parse(string(tmpl_str))
	if tmpl == nil {
		panic("Failed parsing tmpl.")
	}
	check_error(err)

	f, err := os.Create("output/output.html")
	
	check_error(err)
	
	err = tmpl.Execute(f, sound_structs)

	check_error(err)

	err = f.Close()

	check_error(err)

}

func get_sound_names(files []fs.DirEntry) []string {
	re := regexp.MustCompile(`^(.*)\.mp3$`)
	var sound_names []string
	for _, sound := range files {
		file_name := sound.Name()
		sound_name := re.FindStringSubmatch(file_name)[1]
		sound_names = append(sound_names, sound_name)
	}
	return sound_names
}

func check_error(err error) {
	if err != nil {
		panic(err)
	}
}