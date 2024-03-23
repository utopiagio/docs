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
        case "Index":
            return Index
       	default:
			return Index
	}
}

var Index []string = []string{"" + 
"- [**GoWindow**](api.GoWindow#)\n" +
"- [**GoButton**](api.GoButton#)\n" + 
"- [**GoButtonGroup**](api.GoButtonGroup#)\n" + 
"- [**GoCanvas**](api.GoCanvas#)\n" + 
"- [**GoCheckBox**](api.GoCheckBox#)\n" + 
"- [**GoClipboard**](api.GoClipboard#)\n" + 
"- [**GoColor**](api.GoColor#)\n" + 
"- [**GoIconLabel**](api.GoIconLabel#)\n" + 
"- [**GoIconPng**](api.GoIconPng#)\n" + 
"- [**GoIconVg**](api.GoIconVg#)\n" + 
"- [**GoImage**](api.GoImage#)\n" + 
"- [**GoLabel**](api.GoLabel#)\n" + 
"- [**GoListView**](api.GoListView#)\n" + 
"- [**GoListViewItem**](api.GoListViewItem#)\n" + 
"- [**GoLoader**](api.GoLoader#)\n" + 
"- [**GoMenu**](api.GoMenu#)\n" + 
"- [**GoMenuBar**](api.GoMenuBar#)\n" + 
"- [**GoPopupMenu**](api.GoPopupMenu#)\n" + 
"- [**GoPopupWindow**](api.GoPopupWindow#)\n" + 
"- [**GoProgressBar**](api.GoProgressBar#)\n" + 
"- [**GoProgressCircle**](api.GoProgressCircle#)\n" + 
"- [**GoRadioButton**](api.GoRadioButton#)\n" + 
"- [**GoRichText**](api.GoRichText#)\n" + 
"- [**GoScrollBar**](api.GoScrollBar#)\n" + 
"- [**GoSlider**](GgoSlider#)\n" + 
"- [**GoSpacer**](api.GoSpacer#)\n" + 
"- [**GoSwitch**](api.GoSwitch#)\n" + 
"- [**GoTextEdit**](api.GoTextEdit#)\n" + 
"- [**GoTheme**](api.GoTheme#)\n",
"- [**GioObject**](api.GioObject#)\n" + 
"- [**GioWidget**](api.GioWidget#)\n"}

var GoWindow []string = []string{"" + 
"The GoWindow class provides an application window, with a menu bar, and a status bar. [More...](#details)\n" + 
"- [**GoMainWindow**](api.GoWindow#goMainWindow)( windowTitle **string** )  ( hWin [***GoWindowObj**](#goWindowObj) )\n" + 
"- [**GoWindow**](api.GoWindow#goWindow)( windowTitle **string** )  ( hWin [***GoWindowObj**](#goWindowObj) )\n" + 
"- [**GoModalWindow**](api.GoWindow#goModalWindow)( modalStyle **string**, windowTitle **string** )  ( hWin [***GoWindowObj**](#goWindowObj) )\n" + 
"- [**AddPopupMenu**](api.GoWindow#addPopupMenu)()  ( popupMenu [***GoPopupMenuObj**](#goPopupMenuObj) )\n" +
"- [**Centre**](api.GoWindow#centre)()\n" + 
"- [**ClearPopupMenus**](api.GoWindow#clearPopupMenus)**()**\n" + 
"- [**ClientPos**](api.GoWindow#clientPos)**() (** x **int**, y **int )**\n" + 
"- [**Close**](api.GoWindow#close)()\n" + 
"- [**EscFullScreen()**](api.GoWindow#escFullScreen)\n" + 
"- [**GioWindow()**](api.GoWindow#gioWindow) ( gioWin [***app_gio.Window**](#app_gio.Window) )\n" + 
"- [**GoFullScreen**](api.GoWindow#goFullScreen)()\n" + 
"- [**IsMainWindow**](api.GoWindow#isMainWindow)()  ( isMain **bool** )\n" + 
"- [**IsModal**](api.GoWindow#isModal)()  ( ismodal **bool** )\n" + 
"- [**Layout**](api.GoWindow#layout)()  ( layout [***GoLayoutObj**](api.GoLayoutObj#) )\n" + 
"- [**Maximize**](api.GoWindow#maximize)()\n" + 
"- [**Minimize**](api.GoWindow#minimize)()\n" + 
"- [**MenuBar**](api.GoWindow#menubar)()  ( menuBar [***GoMenuBarObj**](#goMenuBarObj]) )\n" + 
"- [**MenuPopup**](api.GoWindow#menuPopup)( idx **int )**  ( popupMenu [***GoPopupMenuObj**](api.GoPopupMenu#) )\n" + 
"- [**ModalStyle**](api.GoWindow#modalStyle)()  ( style **string**)\n" + 
"- [**ObjectType**](api.GoWindow#objectType)()  ( type **string** )\n" + 
"- [**PopupWindow**](api.GoWindow#popupWindow)()  ( popup [***GoPopupWindowObj**](api.GoPopupWindow#) )\n" + 
"- [**Pos**](api.GoWindow#pos)()  ( x **int**, y **int** )\n" + 
"- [**Refresh**](api.GoWindow#refresh)()\n" + 
"- [**SetBorder**](api.GoWindow#setBorder)( style [**GoBorderStyle**](api.GoBorderStyle#), width **int**, radius **int**, color [**GoColor**](api.GoColor#) )\n" + 
"- [**Widget**](api.GoWindow#widget)()  (widget [***GioWidget**](api.GioWidget#) )\n\n" + 


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
"<a name=\"addPopupMenu\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**AddPopupMenu(** popupMenu [***GoPopupMenuObj**](api.GoPopupMenuObj#) **)**\n" + 
"- Add a new popup menu.\n\n" + 
"<a name=\"centre\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**Centre()**\n" + 
"- Centre the window on the client screen.\n\n" + 
"<a name=\"clearPopupMenus\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**ClearPopupMenus()**\n" + 
"- Clear the popup menus.\n\n" + 
"<a name=\"clientPos\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**ClientPos(** x **int,** y **int )**\n" + 
"- Returns the client postion of the window.\nThis is usually set to (0,0)\n\n" + 
"<a name=\"close\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**Close()**\n" + 
"- Close the window.\n\n" + 
"<a name=\"escFullScreen\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**EscFullScreen()**\n" + 
"- Escape fullscreen.\n\n" + 
"<a name=\"goFullScreen\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**GoFullScreen()**\n" + 
"- Switch to fullscreen.\n\n" + 
"<a name=\"isMainWindow\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**IsMainWindow() ( bool )**\n" + 
"- Returns **true** if the window is the main window.\n\n" + 
"<a name=\"isModal\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**IsModal() ( bool )**\n" + 
"- Returns **true** if the window is a modal window.\n\n" + 
"<a name=\"layout\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**Layout() ( layout** [***GoLayoutObj**](api.GoLayout#) **)**\n" + 
"- Returns a pointer to the window centre layout.\n\n" + 
"<a name=\"maximize\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**Maximize()**\n" + 
"- Maximize the window.\n\n" + 
"<a name=\"minimize\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**Minimize()**\n" + 
"- Minimize the window.\n\n" + 
"<a name=\"menubar\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**MenuBar() (** menubar [***GoMenuBarObj**](api.GoMenuBar#) **)**\n" + 
"- Installs and returns a pointer to the window main menu bar.\n\n" + 
"<a name=\"menuPopup\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**MenuPopup() (** popupMenu [***GoPopupMenuObj**](api.GoPopupMenu#) **)**\n" + 
"- Returns a pointer to the popup menu at index idx.\n\n" + 
"<a name=\"modalStyle\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**ModalStyle() (** style **string )**\n" + 
"- Returns the modal style of a modal window.\n\n" + 
"<a name=\"objectType\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**ObjectType() (** type **string )**\n" + 
"- Returns the object type as a string definition \"GoWindowObj\".\n\n" + 
"<a name=\"popupWindow\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**PopupWindow() (** popupWindow [***GoPopupWindowObj**](api.GoPopupWindow#) **)**\n" + 
"- Returns a pointer to the windows popup window.\n\n" + 
"<a name=\"pos\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**Pos(** x **int, y int )**\n" + 
"- Returns the screen postion of the window.\n\n" + 
"<a name=\"refresh\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**Refresh()**\n" + 
"- Refresh the window.\n\n" + 
"<a name=\"setBorder\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**SetBorder(** style **[**GoBorderStyle**](api.GoBorderStyle#), width **int**, radius **int**, color [**GoColor**](api.GoColor#) **)**\n" + 
"- Add a border to the window.\n\n" + 
"<a name=\"widget\"></a> **(ob** [***GoWindowObj)**](api.GoWindow#).**Widget() (** widget [**GioWidget**](api.GioWidget#) **)**\n" + 
"- Returns a pointer to the window widget properties.\n\n"}

var GoButton []string = []string{"" + 
"- **[GoButton](api.GoButton#goButton)(** parent **[GoObject]( GoObject#),** text **string ) (** hObj **[*GoButtonObj](api.GoButton#) )**\n" + 
"- **[Draw](api.GoButton#draw)(** gtx **layout_gio.Context )**\n" + 
"- **[Layout](api.GoButton#layout)(** gtx **layout_gio.Context )**\n" + 
"- **[ObjectType](api.GoWindow#objectType)() (** type **string )**\n" + 
"- **[SetFaceColor](api.GoButton#setFaceColor) (** color **GoColor )**\n" + 
"- **[SetOnClick](api.GoButton#setOnClick) (** f **func )**\n" + 
"- **[SetText](api.GoButton#setText) (** text **string )**\n" + 
"- **[SetTextColor](api.GoButton#setTextColor) (** color **GoColor )**\n" + 
"- **[Text](api.GoButton#text)() (** text **string )**\n" + 
"- **[Widget](api.GoWindow#widget)() (** widget **[*GioWidget](api.GioWidget#) )**\n\n" + 

"##### **Member Function Documentation**\n" + 

"<a name=\"goButton\"></a> **[GoButton](api.GoButton#) ( parent [*GoObject](api.GoObject#), text [string](string) ) ( hObj [*GoButtonObj](api.GoButton#) )**\n" + 
"- Add a new button object, assign a parent and caption text." + 
" Usually a button object will be added to a parent layout.\n\n" + 
" **example:**  button := GoButton(parent, \"Exit\")\n\n</a>" + 

"<a name=\"draw\"></a> **(ob** [***GoButtonObj)**](api.GoButton#).**Draw(gtx [layout_gio.Context](\"https://gioui.org/layout/context.go\")**\n" + 
"- Draw the button object. This function is usually called from a GoLayoutObj.Draw() function" + 
" during the GoWindowObj.render() event and not called directly by a user application.\n" + 
"It is necessary to declare it as a public function to allow new objects using GoObject Interface to be created by the user application.\n\n" + 

"<a name=\"layout\"></a> **(ob** [***GoButtonObj)**](api.GoButton#).**Layout(gtx [layout_gio.Context](\"https://gioui.org/layout/context.go\") )**\n" + 
"- Layout the button object. This function is usually called from a GoButtonObj.Draw() function" + 
" during the GoWindowObj.render() event and not called directly by a user application.\n" + 
"It is necessary to declare it as a public function to facilitate drawing of new objects created by the user application.\n\n" + 

"<a name=\"objectType\"></a> **(ob [*GoButtonObj)](api.GoButton#).ObjectType() ( type [string](string) )**\n" + 
"- Returns the object type as a string definition \"GoButtonObj\".\n\n" + 

"<a name=\"setFaceColor\"></a> **(ob [*GoButtonObj)](api.GoButton#).SetFaceColor( color [GoColor](goColor) )**\n" + 
"- Set the button face color.\n\n" + 

"<a name=\"setOnClick\"></a> **(ob [*GoButtonObj)](api.GoButton#).SetOnClick( f [func](func) )**\n" + 
"- Set the button click function.\n\n" + 

"<a name=\"setText\"></a> **(ob [*GoButtonObj)](api.GoButton#).SetText( text [string](string) )**\n" + 
"- Set the caption text of the button.\n\n" + 

"<a name=\"setTextColor\"></a> **(ob [*GoButtonObj)](api.GoButton#).SetTextColor( color [GoColor](goColor) )**\n" + 
"- Set the text color of the button caption.\n\n" + 

"<a name=\"text\"></a> **(ob [*GoButtonObj)](api.GoButton#).Text() ( text [string](string) )**\n" + 
"- Returns the button caption text.\n\n" + 

"<a name=\"widget\"></a> **(ob [*GoButtonObj)](api.GoButton#).Widget() ( widget [*GioWidget](api.GioWidget#) )**\n" + 
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
"- **[Click](api.GioWidget#click) (** e **pointer_gio.Event )**\n" + 
"- **[Clicked](api.GioWidget#clicked)() (** clicked **bool )**\n" + 
"- **[ClearFocus](api.GioWidget#clearFocus)() (** success **bool )**\n" + 
"- **[HasFocus](api.GioWidget#hasFocus)() (** focus **bool )**\n" + 
"- **[Hide](api.GioWidget#hide)()**\n" + 
"- **[IsFocusEnabled](api.GioWidget#isFocusEnabled)() (** enabled **bool )**\n" + 
"- **[IsHovered](api.GioWidget#isHovered)() (** hovered **bool )**\n" + 
"- **[IsSelected](api.GioWidget#isSelected)() (** selected **bool )**\n" + 
"- **[IsVisible](api.GioWidget#isVisible)() (** visible **bool )**\n" +
"- **[Margin](api.GioWidget#margin)() (** margin [**GoMargin**](#goMargin) **)**\n" + 
"- **[Padding](api.GioWidget#padding)() (** padding [**GoPadding**](#goPadding) **)**\n" + 
"- **[SetBackgroundColor](api.GioWidget#setBackgroundColor)(** color [**GoColor**](api.GoGolor#) **)**\n" + 
"- **[SetBorder](api.GioWidget#setBorder)(** style [**GoBorderStyle**](api.GoBorderStyle#), width **int**, radius **int**, color [**GoColor**](api.GoGolor#) **)**\n" + 
"- **[SetBorderColor](api.GioWidget#setBorderColor)(** color [**GoColor**](api.GoGolor#) **)**\n" + 
"- **[SetBorderRadius](api.GioWidget#setBorderRadius)(** radius **int )**\n" + 
"- **[SetBorderStyle](api.GioWidget#setBorderStyle)(** style [**GoBorderStyle**](api.GoBorderStyle#) **)**\n" + 
"- **[SetBorderWidth](api.GioWidget#setBorderWidth)(** width **int )**\n" + 
"- **[SetFocus](api.GioWidget#setFocus)() (** success **bool )**\n" + 
"- **[ChangeFocus](api.GioWidget#changeFocus)() (** focus **bool )**\n" + 
"- **[SetHeight](api.GioWidget#setHeight)(** height **int )**\n" + 

"- **[SetWidth](GioWidget#setWidth)(** width **int )**\n" + 

"##### **Detail Description** <a name=\"details\"></a>\n" + 


"The GioWidget is the structure of the objects user interface. It receives mouse, keyboard and other events from the operating system, and determines how it is represented on the screen. " + 
" A GioWidget is clipped by its parent, usually a layout widget.\n\n" + 
"If you want to use an object to hold child objects you will probably want to add a layout widget to the parent GioObject. (See Layouts.)\n\n" + 
"The GioWidget structure is embedded by other user controls to allow them to access all the GioWidget properties.\n\n" + 
"Importantly mouse and gesture events are activated by the underlying UtopiaGio framework, when the user control allocates a function to the widgets event.\n" + 
"Key press and release events are activated by the underlying UtopiaGio framework, when the user control sets key filters and allocates a function to either of the widgets key press or release event.\n" + 

""}