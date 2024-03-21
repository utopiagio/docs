// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/docs/apireference.go */

package apireference

func Page(doc string) (content []string) {
	switch doc {
		case "GoMainWindow":
            return GoWindow
        case "GoWindow":
            return GoWindow
        case "GoButton":
            return GoButton
        case "GoButtonGroup":
            return GoButtonGroup
        case "GoCanvas":
            return GoCanvas
        case "GoCheckBox":
            return GoCheckBox
        case "GoImage":
            return GoImage
        case "GoLabel":
            return GoLabel
        case "GoLayout":
            return GoLayout
        case "GoList":
            return GoList
        case "GoMenu":
            return GoMenu
        case "GoMenuBar":
            return GoMenuBar
        case "GoMenuItem":
            return GoMenuItem
        case "GoRadioButton":
            return GoRadioButton
        case "GoRichText":
            return GoRichText
        case "GoSlider":
            return GoSlider
        case "GoSpacer":
            return GoSpacer
        case "GioWidget":
            return GioWidget
       	default:
			return Index
	}
}

var Index []string = []string{"" + 
"- [**GoWindow**](GoWindow#)\n" +
"- [**GoButton**](GoButton#)\n" + 
"- [**GoButtonGroup**](GoButtonGroup#)\n" + 
"- [**GoCanvas**](GoCanvas#)\n" + 
"- [**GoCheckBox**](GoCheckBox#)\n" + 
"- [**GoClipboard**](GoClipboard#)\n" + 
"- [**GoColor**](GoColor#)\n" + 
"- [**GoIconLabel**](GoIconLabel#)\n" + 
"- [**GoIconPng**](GoIconPng#)\n" + 
"- [**GoIconVg**](GoIconVg#)\n" + 
"- [**GoImage**](GoImage#)\n" + 
"- [**GoLabel**](GoLabel#)\n" + 
"- [**GoListView**](GoListView#)\n" + 
"- [**GoListViewItem**](GoListViewItem#)\n" + 
"- [**GoLoader**](GoLoader#)\n" + 
"- [**GoMenu**](GoMenu#)\n" + 
"- [**GoMenuBar**](GoMenuBar#)\n" + 
"- [**GoPopupMenu**](GoPopupMenu#)\n" + 
"- [**GoPopupWindow**](GoPopupWindow#)\n" + 
"- [**GoProgressBar**](GoProgressBar#)\n" + 
"- [**GoProgressCircle**](GoProgressCircle#)\n" + 
"- [**GoRadioButton**](GoRadioButton#)\n" + 
"- [**GoRichText**](GoRichText#)\n" + 
"- [**GoScrollBar**](GoScrollBar#)\n" + 
"- [**GoSlider**](GgoSlider#)\n" + 
"- [**GoSpacer**](GoSpacer#)\n" + 
"- [**GoSwitch**](GoSwitch#)\n" + 
"- [**GoTextEdit**](GoTextEdit#)\n" + 
"- [**GoTheme**](GoTheme#)\n",
"- [**GioObject**](GioObject#)\n" + 
"- [**GioWidget**](GioWidget#)\n"}

var GoWindow []string = []string{"" + 
"The GoWindow class provides an application window, with a menu bar, and a status bar. [More...](#details)\n" + 
"- [**GoMainWindow**](GoWindow#goMainWindow)( windowTitle **string** )  ( hWin [***GoWindowObj**](#goWindowObj) )\n" + 
"- [**GoWindow**](GoWindow#goWindow)( windowTitle **string** )  ( hWin [***GoWindowObj**](#goWindowObj) )\n" + 
"- [**GoModalWindow**](GoWindow#goModalWindow)( modalStyle **string**, windowTitle **string** )  ( hWin [***GoWindowObj**](#goWindowObj) )\n" + 
"- [**AddPopupMenu**](GoWindow#addPopupMenu)()  ( popupMenu [***GoPopupMenuObj**](#goPopupMenuObj) )\n" +
"- [**Centre**](GoWindow#centre)()\n" + 
"- [**ClearPopupMenus**](GoWindow#clearPopupMenus)**()**\n" + 
"- [**ClientPos**](GoWindow#clientPos)**() (** x **int**, y **int )**\n" + 
"- [**Close**](GoWindow#close)()\n" + 
"- [**EscFullScreen()**](GoWindow#escFullScreen)\n" + 
"- [**GioWindow()**](GoWindow#gioWindow) ( gioWin [***app_gio.Window**](#app_gio.Window) )\n" + 
"- [**GoFullScreen**](GoWindow#goFullScreen)()\n" + 
"- [**IsMainWindow**](GoWindow#isMainWindow)()  ( isMain **bool** )\n" + 
"- [**IsModal**](GoWindow#isModal)()  ( ismodal **bool** )\n" + 
"- [**Layout**](GoWindow#layout)()  ( layout [***GoLayoutObj**](GoLayoutObj#) )\n" + 
"- [**Maximize**](GoWindow#maximize)()\n" + 
"- [**Minimize**](GoWindow#minimize)()\n" + 
"- [**MenuBar**](GoWindow#menubar)()  ( menuBar [***GoMenuBarObj**](#goMenuBarObj]) )\n" + 
"- [**MenuPopup**](GoWindow#menuPopup)( idx **int )**  ( popupMenu [***GoPopupMenuObj**](GoPopupMenu#) )\n" + 
"- [**ModalStyle**](GoWindow#modalStyle)()  ( style **string**)\n" + 
"- [**ObjectType**](GoWindow#objectType)()  ( type **string** )\n" + 
"- [**PopupWindow**](GoWindow#popupWindow)()  ( popup [***GoPopupWindowObj**](GoPopupWindow#) )\n" + 
"- [**Pos**](GoWindow#pos)()  ( x **int**, y **int** )\n" + 
"- [**Refresh**](GoWindow#refresh)()\n" + 
"- [**SetBorder**](GoWindow#setBorder)( style [**GoBorderStyle**](GoBorderStyle#), width **int**, radius **int**, color [**GoColor**](GoColor#) )\n" + 
"- [**Widget**](GoWindow#widget)()  (widget [***GioWidget**](GioWidget#) )\n\n" + 


"##### **Detail Description** <a name=\"details\"></a>\n" + 
"The GoWindow class provides an application window, with a menu bar, and a status bar.\n\n" + 
"The application main window is always the first GoWindow created by the application. Further windows will be created as floating popup windows.\n" + 
"If the main window is closed, the application will be terminated and all other windows destroyed.\n\n" + 
"All windows may provide menus and a status bar around a central layout, containing application specific features such as a text editor, drawing canvas, or other communications features.\n\n" + 
"The main window can not be created before an application instance is declared. See the example below:\n\n" + 
"```\n" + 
"    func main() {\n" + 
"        // create application instance before any other objects\n" + 
"        app = ui.GoApplication(\"APIViewer\")\n" + 
"        // create application window\n" + 
"        mainwin = ui.GoMainWindow(\"API Viewer - UtopiaGio Package\")\n" + 
"        // set the window layout style to stack widgets vertically\n" + 
"        mainwin.SetLayoutStyle(ui.VFlexBoxLayout)\n" + 
"        mainwin.SetMargin(10,10,10,10)\n" + 
"        mainwin.SetBorder(ui.BorderSingleLine, 2, 10, ui.Color_Blue)\n" + 
"        mainwin.SetPadding(10,10,10,10)\n" + 
"        mainwin.SetSize(1000, 600)\n" + 
"        mainwin.SetPos(100,100)\n" + 
"        // show the main window\n" + 
"        mainwin.Show()\n" + 
"        // run the application\n" + 
"        app.Run()\n" + 
"    }\n" + 
"```\n" + 
"The menubar is activated by declaring and showing the menubar.\n" + 
"```\n" + 
"   // setup the main window MenuBar\n" + 
"    menuBar := mainwin.MenuBar()\n" + 
"    menuBar.Show() {\n" + 
"```\n" + 
"and adding menus to the menubar and menu items to the menus.\n" + 
"```\n" + 
"    mnuDocs := menuBar.AddMenu(\"Documentation\")\n" + 
"    mnuOverview := mnuDocs.AddItem(\"Overview\")\n" + 
"```\n" + 
"The AddMenu function creates a new menu and in turn places it in the menubar and " + 
"the AddItem function creates a new menu item and in turn places it in to the menu.\n" + 
"The item has to associated with an action (a function to be called when the menu item is clicked).\n" + 
"```\n" + 
"    mnuOverview.SetOnClick(ActionOverview_Clicked)\n" + 
"```\n" + 
"It is perhaps easier to call the AddAction function which creates a new menu item and adds the action function all in one call.\n" + 
"```\n" + 
"    mnuDocs := menuBar.AddMenu(\"Documentation\")\n" + 
"    mnuOverview := mnuDocs.AddAction(\"Overview\", ActionOverview_Clicked)\n" + 
"```\n" + 

"##### **Member Function Documentation**\n" + 
"<a name=\"addPopupMenu\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**AddPopupMenu(** popupMenu [***GoPopupMenuObj**](GoPopupMenuObj#) **)**\n" + 
"- Add a new popup menu.\n\n" + 
"<a name=\"centre\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**Centre()**\n" + 
"- Centre the window on the client screen.\n\n" + 
"<a name=\"clearPopupMenus\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**ClearPopupMenus()**\n" + 
"- Clear the popup menus.\n\n" + 
"<a name=\"clientPos\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**ClientPos(** x **int,** y **int )**\n" + 
"- Returns the client postion of the window.\nThis is usually set to (0,0)\n\n" + 
"<a name=\"close\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**Close()**\n" + 
"- Close the window.\n\n" + 
"<a name=\"escFullScreen\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**EscFullScreen()**\n" + 
"- Escape fullscreen.\n\n" + 
"<a name=\"goFullScreen\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**GoFullScreen()**\n" + 
"- Switch to fullscreen.\n\n" + 
"<a name=\"isMainWindow\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**IsMainWindow() ( bool )**\n" + 
"- Returns **true** if the window is the main window.\n\n" + 
"<a name=\"isModal\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**IsModal() ( bool )**\n" + 
"- Returns **true** if the window is a modal window.\n\n" + 
"<a name=\"layout\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**Layout() ( layout** [***GoLayoutObj**](GoLayout#) **)**\n" + 
"- Returns a pointer to the window centre layout.\n\n" + 
"<a name=\"maximize\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**Maximize()**\n" + 
"- Maximize the window.\n\n" + 
"<a name=\"minimize\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**Minimize()**\n" + 
"- Minimize the window.\n\n" + 
"<a name=\"menubar\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**MenuBar() (** menubar [***GoMenuBarObj**](GoMenuBar#) **)**\n" + 
"- Installs and returns a pointer to the window main menu bar.\n\n" + 
"<a name=\"menuPopup\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**MenuPopup() (** popupMenu [***GoPopupMenuObj**](GoPopupMenu#) **)**\n" + 
"- Returns a pointer to the popup menu at index idx.\n\n" + 
"<a name=\"modalStyle\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**ModalStyle() (** style **string )**\n" + 
"- Returns the modal style of a modal window.\n\n" + 
"<a name=\"objectType\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**ObjectType() (** type **string )**\n" + 
"- Returns the object type as a string definition \"GoWindowObj\".\n\n" + 
"<a name=\"popupWindow\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**PopupWindow() (** popupWindow [***GoPopupWindowObj**](GoPopupWindow#) **)**\n" + 
"- Returns a pointer to the windows popup window.\n\n" + 
"<a name=\"pos\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**Pos(** x **int, y int )**\n" + 
"- Returns the screen postion of the window.\n\n" + 
"<a name=\"refresh\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**Refresh()**\n" + 
"- Refresh the window.\n\n" + 
"<a name=\"setBorder\"></a> **(ob** [***GoWindowObj**](GoWindow#)**)**.**SetBorder(** style **[**GoBorderStyle**](GoBorderStyle#), width **int**, radius **int**, color [**GoColor**](GoColor#) **)**\n" + 
"- Add a border to the window.\n\n" + 
"<a name=\"widget\"></a> **(ob** [***GoWindowObj)**](GoWindow#).**Widget() (** widget [**GioWidget**](GioWidget#) **)**\n" + 
"- Returns a pointer to the window widget properties.\n\n"}

var GoButton []string = []string{"" + 
"- **[GoButton](GoButton#goButton)(** parent **[GoObject]( GoObject#),** text **string ) (** hObj **[*GoButtonObj](GoButton#) )**\n" + 
"- **[Draw](GoButton#draw)(** gtx **layout_gio.Context )**\n" + 
"- **[Layout](GoButton#layout)(** gtx **layout_gio.Context )**\n" + 
"- **[ObjectType](GoWindow#objectType)() (** type **string )**\n" + 
"- **[SetFaceColor](GoButton#setFaceColor) (** color **GoColor )**\n" + 
"- **[SetOnClick](GoButton#setOnClick) (** f **func )**\n" + 
"- **[SetText](GoButton#setText) (** text **string )**\n" + 
"- **[SetTextColor](GoButton#setTextColor) (** color **GoColor )**\n" + 
"- **[Text](GoButton#text)() (** text **string )**\n" + 
"- **[Widget](GoWindow#widget)() (** widget **[*GioWidget](GioWidget#) )**\n\n" + 

"##### **Member Function Documentation**\n" + 

"<a name=\"goButton\"></a> **[GoButton](GoButton#) ( parent [*GoObject](GoObject#), text [string](string) ) ( hObj [*GoButtonObj](GoButton#) )**\n" + 
"- Add a new button object, assign a parent and caption text." + 
" Usually a button object will be added to a parent layout.\n\n" + 
" **example:**  button := GoButton(parent, \"Exit\")\n\n</a>" + 

"<a name=\"draw\"></a> **(ob** [***GoButtonObj)**](GoButton#).**Draw(gtx [layout_gio.Context](\"https://gioui.org/layout/context.go\")**\n" + 
"- Draw the button object. This function is usually called from a GoLayoutObj.Draw() function" + 
" during the GoWindowObj.render() event and not called directly by a user application.\n" + 
"It is necessary to declare it as a public function to allow new objects using GoObject Interface to be created by the user application.\n\n" + 

"<a name=\"layout\"></a> **(ob** [***GoButtonObj)**](GoButton#).**Layout(gtx [layout_gio.Context](\"https://gioui.org/layout/context.go\") )**\n" + 
"- Layout the button object. This function is usually called from a GoButtonObj.Draw() function" + 
" during the GoWindowObj.render() event and not called directly by a user application.\n" + 
"It is necessary to declare it as a public function to facilitate drawing of new objects created by the user application.\n\n" + 

"<a name=\"objectType\"></a> **(ob [*GoButtonObj)](GoButton#).ObjectType() ( type [string](string) )**\n" + 
"- Returns the object type as a string definition \"GoButtonObj\".\n\n" + 

"<a name=\"setFaceColor\"></a> **(ob [*GoButtonObj)](GoButton#).SetFaceColor( color [GoColor](goColor) )**\n" + 
"- Set the button face color.\n\n" + 

"<a name=\"setOnClick\"></a> **(ob [*GoButtonObj)](GoButton#).SetOnClick( f [func](func) )**\n" + 
"- Set the button click function.\n\n" + 

"<a name=\"setText\"></a> **(ob [*GoButtonObj)](GoButton#).SetText( text [string](string) )**\n" + 
"- Set the caption text of the button.\n\n" + 

"<a name=\"setTextColor\"></a> **(ob [*GoButtonObj)](GoButton#).SetTextColor( color [GoColor](goColor) )**\n" + 
"- Set the text color of the button caption.\n\n" + 

"<a name=\"text\"></a> **(ob [*GoButtonObj)](GoButton#).Text() ( text [string](string) )**\n" + 
"- Returns the button caption text.\n\n" + 

"<a name=\"widget\"></a> **(ob [*GoButtonObj)](GoButton#).Widget() ( widget [*GioWidget](GioWidget#) )**\n" + 
"- Returns a pointer to the button widget properties.\n\n"}


var GoButtonGroup []string = []string{""}
var GoCanvas []string = []string{""}
var GoCheckBox []string = []string{""}
var GoImage []string = []string{""}
var GoLabel []string = []string{""}
var GoLayout []string = []string{""}
var GoList []string = []string{""}
var GoMenu []string = []string{""}
var GoMenuBar []string = []string{""}
var GoMenuItem []string = []string{""}
var GoRadioButton []string = []string{""}
var GoRichText []string = []string{""}
var GoSlider []string = []string{""}
var GoSpacer []string = []string{""}
var GioObject []string = []string{""} 
var GioWidget []string = []string{"" + 
"The GioWidgetObj is the base widget object for all user interface objects. [More...](#details)\n" + 
"- **[Click](GioWidget#click) (** e **pointer_gio.Event )**\n" + 
"- **[Clicked](GioWidget#clicked)() (** clicked **bool )**\n" + 
"- **[ClearFocus](GioWidget#clearFocus)() (** success **bool )**\n" + 
"- **[HasFocus](GioWidget#hasFocus)() (** focus **bool )**\n" + 
"- **[Hide](GioWidget#hide)()**\n" + 
"- **[IsFocusEnabled](GioWidget#isFocusEnabled)() (** enabled **bool )**\n" + 
"- **[IsHovered](GioWidget#isHovered)() (** hovered **bool )**\n" + 
"- **[IsSelected](GioWidget#isSelected)() (** selected **bool )**\n" + 
"- **[IsVisible](GioWidget#isVisible)() (** visible **bool )**\n" +
"- **[Margin](GioWidget#margin)() (** margin [**GoMargin**](#goMargin) **)**\n" + 
"- **[Padding](GioWidget#padding)() (** padding [**GoPadding**](#goPadding) **)**\n" + 
"- **[SetBackgroundColor](GioWidget#setBackgroundColor)(** color [**GoColor**](GoGolor#) **)**\n" + 
"- **[SetBorder](GioWidget#setBorder)(** style [**GoBorderStyle**](GoBorderStyle#), width **int**, radius **int**, color [**GoColor**](GoGolor#) **)**\n" + 
"- **[SetBorderColor](GioWidget#setBorderColor)(** color [**GoColor**](GoGolor#) **)**\n" + 
"- **[SetBorderRadius](GioWidget#setBorderRadius)(** radius **int )**\n" + 
"- **[SetBorderStyle](GioWidget#setBorderStyle)(** style [**GoBorderStyle**](GoBorderStyle#) **)**\n" + 
"- **[SetBorderWidth](GioWidget#setBorderWidth)(** width **int )**\n" + 
"- **[SetFocus](GioWidget#setFocus)() (** success **bool )**\n" + 
"- **[ChangeFocus](GioWidget#changeFocus)() (** focus **bool )**\n" + 
"- **[SetHeight](GioWidget#setHeight)(** height **int )**\n" + 

"- **[SetWidth](GioWidget#setWidth)(** width **int )**\n" + 

"##### **Detail Description** <a name=\"details\"></a>\n" + 


"The GioWidget is the structure of the objects user interface. It receives mouse, keyboard and other events from the operating system, and determines how it is represented on the screen. " + 
" A GioWidget is clipped by its parent, usually a layout widget.\n\n" + 
"If you want to use an object to hold child objects you will probably want to add a layout widget to the parent GioObject. (See Layouts.)\n\n" + 
"The GioWidget structure is embedded by other user controls to allow them to access all the GioWidget properties.\n\n" + 
"Importantly mouse and gesture events are activated by the underlying UtopiaGio framework, when the user control allocates a function to the widgets event.\n" + 
"Key press and release events are activated by the underlying UtopiaGio framework, when the user control sets key filters and allocates a function to either of the widgets key press or release event.\n" + 

""}