// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/docs/overview.go */

package overview

func Page(doc string) (content []string) {
	switch doc {
		case "Overview":
			return Overview
		default:
			return Overview
	}
}

var Overview []string = []string{"[**UtopiaGio**](https://github.com/utopiagio/utopia) is a Go framework library module calling the [**Gio**](https://gioui.org) library module. Gio is a cross-platform immediate mode GUI.\n\n" + 
"The [**GoApplication**](api.GoApplication#) object maintains a list of GoWindows and manages the control of the GoWindows and their running threads.\n\n" +
"Each [**GoWindow**](api.GoWindow#) runs it's own message loop, but it will be possible to send and receive communications over channels between windows.\n\n" +
"The framework allows the building of more complex programs without the necessity to write code calling the Gio module directely. In turn this means reduced calls to Gio directly, but with the ability to still write specific Gio routines. It is also possible to use all of the Gio widgets by encapsulating within a new UtopiaGio widget structure and calling the Layout function from the GoObject.Draw() function.\n\n" +
"Inheritance is achieved by embedding the new [**GioObject**](api.GioObject#), and access to the user interface is provided by embedding the new [**GioWidget**](api.GioWidget#). " +
"New layout methods have been introduced requiring a very small change to the Gio layout package. The Gio widget package is still used with some of the widgets, but the intention is to move any relevant code for GoWidgets to the **utopia/internal/widget** package.\n\n" +
"Access to the underlying OS Screen and Main Window is provided through the **utopia/desktop** package, with the possibily to retrieve position, size and scaling of gio windows. The [**Pos**](api.GoWindow#pos) function has been added to the Gio package, which along with the [**Size**](api.GoWindow#size) function allows positioning and sizing of the gio window. Also available at run time using the utopia.GoWindowObj [**SetPos()**](api.GoWindow#setPos) and [**SetSize()**](api.GoWindow#setSize) functions.\n\n" + 

"###### **A simple GoMainWindow**\n" +
"```\n" +
"   package main\n" +
"\n" +
"   import (\n" +
"       ui \"github.com/utopiagio/utopia\"\n" +
"   )\n" +
"\n" +
"   var mainwin *ui.GoWindowObj\n" +
"\n" +
"   func main() {\n" +
"       // create application instance before any other objects\n" +
"       app := ui.GoApplication(\"GoMainWindowDemo\")\n" +
"\n" +	
"       // create application window\n" +
"       mainwin = ui.GoMainWindow(\"GoMainWindow Demo - UtopiaGio Package\")\n" +
"\n" +	
"       // set the window layout style to stack widgets vertically\n" +
"       mainwin.SetLayoutStyle(ui.VFlexBoxLayout)\n" +
"       mainwin.SetMargin(10,10,10,10)\n" +
"       mainwin.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)\n" +
"       mainwin.SetPadding(10,10,10,10)\n" +
"\n" +	
"       // show the application window\n" +
"       mainwin.Show()\n" +
"\n" +
"       // run the application\n" +
"       app.Run()\n" +
"   }\n" +
"```\n\n" +

"Every GoWindowObj uses a main layout to allow the positioning of child controls (GioWidgets). The main layout is accessible through the GoWindowObj.Layout() function.\n" + 
"All child controls are created by passing the parent control as the first parameter.\n" +
"```\n" +
"    lblHello := ui.GoLabel(mainwin.Layout(), \"Hello\")\n" +
"```\n" +
"###### **GoLayout**\n" +
"Usually the main window will be constructed using multiple layouts to allow the positioning of child controls into regions within the main window.\n" +
"```\n" +
"    ....\n" +
"    layoutTop := ui.GoHFlexBoxLayout(mainwin.Layout())\n" +
"    layoutBottom := ui.GoHFlexBoxLayout(mainwin.Layout())\n" +
"    ....\n" +
"```\n\n" +
"GoLayout has four possible configurations.\n\n" + 
// +
"1 **GoHBoxLayout** aligns the child controls horizontally within the layout. No attempt is made to constrain the child controls to fit within the bounds of the layout control. However if the chid controls exceed the horizontal extents of the layout a horizontal scroll bar will be displayed and the layout can be scrolled.\n\n" +

"2 **GoVBoxLayout** aligns the child controls vertically within the layout. No attempt is made to constrain the child controls to fit within the bounds of the layout control. However if the chid controls exceed the vertical extents of the layout a vertical scroll bar will be displayed and the layout can be scrolled.\n\n" +

"3 **GoHFlexBoxLayout** aligns the child controls horizontally and attempts to constrain the child control widths to fit within the layout.\n\n" +

"4 **GoVFlexBoxLayout** aligns the child controls vertically and attempts to constrain the child control heights to fit within the layout.\n\n" +

"The sizing method of each child control is determined by the child control's **GoSizePolicy**.\n" +

"All controls (GioWidgets), including layouts, have margin, border and padding (GoMargin ,GoBorder, GoPadding) properties, which can be set at run time. Defaults are provided for controls where possible.\n\n" +
"These are the main windows layout properties\n" +
"```\n" +
"   ....\n" +
"   mainwin.SetMargin(10,10,10,10)\n" +
"   mainwin.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)\n" +
"   mainwin.SetPadding(10,10,10,10)\n" +
"   ....\n" +
"```\n" +
"Usually the main window will only have padding.\n" +
"```\n" +
"   mainwin.SetPadding(10,10,10,10)\n" +
"```\n" +
"Then to add child layouts to the main window and set parameters\n" +
"```\n" +
"   ....\n" +
"   layoutTop := ui.GoHFlexBoxLayout(mainwin.Layout())\n" +
"   layoutTop.SetMargin(0,0,0,0)                          // Same as default layout margin\n" +
"   layoutTop.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)\n" +
"   layoutTop.SetPadding(10,10,10,10)\n" +
"\n" +
"   ui.GoSpacer(win.Layout(), 10)                         // Add spacer in between layouts\n" +
"\n" +
"   layoutBottom := ui.GoHFlexBoxLayout(mainwin.Layout())\n" +
"   layoutBottom.SetSizePolicy(ui.ExpandingWidth, ui.PreferredHeight)    // GoSizePolicy\n" +
"   layoutBottom.SetMargin(0,0,0,0)                       // Same as default layout margin\n" +
"   layoutBottom.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)\n" +
"   layoutBottom.SetPadding(0,0,0,0)\n" +
"   ....\n" +
"```\n" +
"###### **GoSizePolicy**\n" +
"GioWidget sizing can be controlled using a sizing policy. There are basically six settings available **FixedWidth** and **FixedHeight**, **PreferredWidth** and **PreferredHeight**, **ExpandingWidth** and **ExpandingHeight**.\n\n" +

"1 **Fixed**  restrains the widget to its Width and Height parameters.\n\n" +

"2 **Preferred**  restrains the widget to the dimensions of its children.\n\n" +

"3 **Expanding**  expands the widget to use all the available remaining space.\n\n" +

"The default for **layouts** is ExpandingWidth, ExpandingHeight.\n" +
"###### **GoTextEdit**\n" +
"Add a **GoTextEdit** control to the top layout and set its SizePolicy and Font.\n" +
"```\n" +
"   txtPad := ui.GoTextEdit(layoutTop, \"Enter text here.\")\n" +
"   txtPad.SetSizePolicy(ui.ExpandingWidth, ui.ExpandingHeight)\n" +
"   txtPad.SetFont(\"Go\", ui.Regular, ui.Bold)\n" +
"```\n" +
"Notice the parent object **layoutTop**. This declaration renders the TextEdit as a child of this layout.\n" +
"###### **GoButton**\n" +
"Add a **GoButtonObj** to the bottom layout and set its border and padding along with the onClick action function.\n" +
"```\n" +
"   btnClose := ui.GoButton(layoutBottom, \"Close\")\n" +
"   btnClose.SetBorder(ui.BorderSingleLine, 1, 6, ui.Color_Blue)\n" +
"   btnClose.SetPadding(4,4,4,4)\n" +
"   btnClose.SetOnClick(ActionExit_Clicked)\n" +
"```\n\n" +
"Notice the parent object **layoutBottom**. This declaration renders the Button as a child of this layout. Also because the layout has a SizePolicy of ui.PreferredHeight, the layout will size to contain the button object and not expand.\n\n" +

"The **GoButtonObj** also has the default SizePolicy of ui.PreferredWidth and ui.PreferredHeight resulting in a button just big enough to display the caption of the button plus some default padding.\n\n" +

"The function ActionExit_Clicked() must be declared outside the package main() function as an external function.\n" +
"```\n" +
"    func ActionExit_Clicked() {\n" +
"       log.Println(\"ActionExit_Clicked().......\")\n" +
"       os.Exit(0)\n" +
"    }\n" +
"```\n" +

"###### **To see a demo GoHello run:**\n" +
"```\n" +
"   go run github.com/utopiagio/demos/GoHello@latest\n" +
"```\n" +

""}