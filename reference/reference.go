// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/utopia/docs/apireference.go */

package reference

func Page(doc string) (content []string) {
	switch doc {
		case "Demo_GoHello":
            return Demo_GoHello
        default:
			return Index
	}
}

var Index []string = []string{"" + 
	"- [**Demo_GoHello**](ref.Demo_GoHello#)\n"}

var Demo_GoHello []string = []string{"" + 
	"### GoHello Example\n"}

/*"### GoButton\n" + 
"Add a **GoButtonObj** to the bottom layout and set its border and padding along with the onClick action function.\n" + 
"```\n" + 
"    btnClose := ui.GoButton(layoutBottom, \"Close\")\n" + 
"    btnClose.SetBorder(ui.BorderSingleLine, 1, 6, ui.Color_Blue)\n" + 
"    btnClose.SetPadding(4,4,4,4)\n" + 
"    btnClose.SetOnClick(ActionExit_Clicked)\n" + 
"```\n" + 
"Notice the parent object **layoutBottom**. This declaration renders the Button as a child of this layout. Also because the layout has a SizePolicy of ui.PreferredHeight, the layout will size to contain the button object and not expand.\n\n" + 

"The **GoButtonObj** also has the default SizePolicy of ui.PreferredWidth and ui.PreferredWidth resulting in a button just big enough to display the caption of the button plus some default padding.\n\n" + 

"The function ActionExit_Clicked() must be declared outside the package main() function as an external function.\n\n" + 
"```\n" + 
"    func ActionExit_Clicked() {\n" + 
"        log.Println(\"ActionExit_Clicked().......\")\n" + 
"        os.Exit(0)\n" + 
"    }\n" + 
"```"*/