// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/docs/apireference.go */

package apireference

func Page(doc string) (content []string) {
	switch doc {
        case "GoApplication":
            return GoApplication
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
        case "GoListView":
            return GoListView
        case "GoListViewItem":
            return GoListViewItem
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
        case "GioObject":
            return GioObject    
        case "GioWidget":
            return GioWidget
        case "Index":
            return Index
       	default:
			return PageNotFound
	}
}

var PageNotFound []string = []string{"" + 
"## **Page Not Found.**\n\n" + 
"###### **The page you are searching for could not be found.**\n" + 
"###### **The reference manual is in the process of being created!**\n" + 
"###### **Sorry for the inconvenience.**\n" + 
"\n"}

var Index []string = []string{"" + 
"- [**GoApplication**](api.GoApplication#)\n" +
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

var GoApplication []string = []string{"" + 
"The GoApplication maintains a list of GoWindows and manages the control of all the GoWindowObjs and their running threads. [More...](#details)\n\n" + 

"[**GoApplication**](api.GoApplication#goApplication)( appName [string]())  ( app [***GoApplicationObj**](#properties) )\n\n" + 
"- [**AddWindow**](api.GoApplication#addWindow)( win [***GoWindowObj**](api.GoWindow#properties) )\n" +
"- [**ClipBoard**](api.GoApplication#clipBoard)()  ( clipBoard [***GoClipBoardObj**](api.GoClipBoard#properties) )\n" +
"- [**RemoveWindow**](api.GoApplication#removeWindow)( win [***GoWindowObj**](api.GoWindow#properties) )\n" +
"- [**Run**](api.GoApplication#run)()\n" + 
"- [**SetModal**](api.GoApplication#setModal)()  ( modalWin [***GoWindowObj**](api.GoWindow#properties) )\n" +
"- [**Theme**](api.GoApplication#theme)()  ( theme [***GoThemeObj**](api.GoTheme#properties) )\n" +

"##### **Detail Description** <a name=\"details\"></a>\n" + 

"The GoApplication maintains a list of GoWindows and manages the control of all the GoWindowObjs and their running threads. [More...](#details)\n\n" + 

"There is only one GoApplicationObj in any one GUI application. The GoApplicationObj is accessible through the global pointer GoApp.\n\n" + 

"Since the GoApplicationObj does much of the initialization, it must be created before any other objects.\n\n" + 

"It provides the desktop() and the clipboard() objects.\n\n" +

"###### **GoApplicationObj Properties** <a name=\"properties\"></a>\n\n" + 

"[type]() **GoApplicationObj** [struct](){}\n" + 
"- name [string]()\n" + 
"- windows **[[]*GoWindowObj](api.GoWindow#properties)**\n" + 
"- clipboard [***GoClipBoardObj**](api.GoClipBoard#properties)\n" + 
"- keyboard [***GoKeyboardObj**](api.GoKeyBoard#properties)\n" + 
"- theme [***GoThemeObj**](api.GoTheme#properties)\n" + 
"- shaper [***text_gio.Shaper**]()\n" + 
"- mode [**GoApplicationMode**](api.GoApplicationMode#)\n" + 
"- modalWindow [***GoWindowObj**](api.GoWindow#properties)\n" + 
"\n\n" + 

"##### **Member Function Documentation**\n" + 

"<a name=\"goApplication\"></a> [**GoApplication**](api.GoApplication#goApplication)( appName **string** )  ( app [***GoApplicationObj**](#properties) )\n" + 
"- Initialise the application. Instantiate the GoApp global reference.\n\n" + 
"<a name=\"addWindow\"></a> **(ob** [***GoApplicationObj**](api.GoApplication#)**)**.**AddWindow(** win [***GoWindowObj**](api.GoWindow#properties) **)**\n" + 
"- Add a new window to the application.\n\n" + 
"<a name=\"clipBoard\"></a> **(ob** [***GoApplicationObj**](api.GoApplication#)**)**.**ClipBoard() (** clipboard [***GoClipBoardObj**](api.GoClipBoard#properties) **)**\n" + 
"- Return the application clipboard.\n\n" + 
"<a name=\"removeWindow\"></a> **(ob** [***GoApplicationObj**](api.GoApplication#)**)**.**RemoveWindow(** win [***GoWindowObj**](api.GoWindow#properties) **)**\n" + 
"- Remove a window from the application.\nIf the window is the main window then the application will be shut down.\n\n" + 
"<a name=\"run\"></a> **(ob** [***GoApplicationObj**](api.GoApplication#)**)**.**Run()**\n" + 
"- Run the application main loop.\n\n" + 
"<a name=\"setModal\"></a> **(ob** [***GoApplicationObj**](api.GoApplication#)**)**.**SetModal(** modalWin [***GoWindowObj**](api.GoWindow#) **)**\n" + 
"- Set the window to run as a modal window.\n The application will be set to run in ModalMode and all other windows running under the application will be disabled.\n\n" + 
"<a name=\"theme\"></a> **(ob** [***GoApplicationObj**](api.GoApplication#)**)**.**Theme() (** theme [***GoThemeObj**](api.GoTheme#) **)**\n" + 
"- Return the application main theme.\n\n" + 
"\n"} 


var GoWindow []string = []string{"" + 
"The GoWindow object provides an application window, with a menu bar, and a status bar. [More...](#details)\n\n" + 

"- [**GoMainWindow**](api.GoWindow#goMainWindow)( windowTitle [**string**]() )  ( hWin [***GoWindowObj**](#properties) )\n" + 
"- [**GoWindow**](api.GoWindow#goWindow)( windowTitle [**string**]() )  ( hWin [***GoWindowObj**](#properties) )\n" + 
"- [**GoModalWindow**](api.GoWindow#goModalWindow)( modalStyle [**string**](), windowTitle [**string**]() )  ( hWin [***GoWindowObj**](#properties) )\n\n" + 
"member functions...\n\n" + 
"- [**AddPopupMenu**](api.GoWindow#addPopupMenu)()  ( popupMenu [***GoPopupMenuObj**](#goPopupMenuObj) )\n" +
"- [**Centre**](api.GoWindow#centre)()\n" + 
"- [**ClearPopupMenus**](api.GoWindow#clearPopupMenus)()\n" + 
"- [**ClientPos**](api.GoWindow#clientPos)()  ( x [**int**](), y [**int**]() )\n" + 
"- [**Close**](api.GoWindow#close)()\n" + 
"- [**EscFullScreen**](api.GoWindow#escFullScreen)()\n" + 
"- [**GioWindow**](api.GoWindow#gioWindow)()  ( gioWin [***app_gio.Window**](#app_gio.Window) )\n" + 
"- [**GoFullScreen**](api.GoWindow#goFullScreen)()\n" + 
"- [**IsMainWindow**](api.GoWindow#isMainWindow)()  ( isMain [**bool**]() )\n" + 
"- [**IsModal**](api.GoWindow#isModal)()  ( ismodal [**bool**]() )\n" + 
"- [**Layout**](api.GoWindow#layout)()  ( layout [***GoLayoutObj**](api.GoLayoutObj#) )\n" + 
"- [**Maximize**](api.GoWindow#maximize)()\n" + 
"- [**Minimize**](api.GoWindow#minimize)()\n" + 
"- [**MenuBar**](api.GoWindow#menubar)()  ( menuBar [***GoMenuBarObj**](api.GoMenuBar#]) )\n" + 
"- [**MenuPopup**](api.GoWindow#menuPopup)( idx [**int**]() )  ( popupMenu [***GoPopupMenuObj**](api.GoPopupMenu#) )\n" + 
"- [**ModalStyle**](api.GoWindow#modalStyle)()  ( style [**string**]() )\n" + 
"- [**ObjectType**](api.GoWindow#objectType)()  ( type [**string**]() )\n" + 
"- [**PopupWindow**](api.GoWindow#popupWindow)()  ( popup [***GoPopupWindowObj**](api.GoPopupWindow#) )\n" + 
"- [**Pos**](api.GoWindow#pos)()  ( [**int**](), y [**int**]() )\n" + 
"- [**Refresh**](api.GoWindow#refresh)()\n" + 
"- [**SetBorder**](api.GoWindow#setBorder)( style [**GoBorderStyle**](api.GoBorderStyle#), width [**int**](), radius [**int**](), color [**GoColor**](api.GoColor#) )\n" + 
"- [**SetBorderColor**](api.GoWindow#setBorderColor)( color [**GoColor**](api.GoColor#) )\n" + 
"- [**SetBorderRadius**](api.GoWindow#setBorderRadius)( radius [**int**]() )\n" + 
"- [**SetBorderStyle**](api.GoWindow#setBorderStyle)( style [**GoBorderStyle**](api.GoBorderStyle#) )\n" + 
"- [**SetBorderWidth**](api.GoWindow#setBorderWidth)( width [**int**]() )\n" + 
"- [**SetLayoutStyle**](api.GoWindow#setLayoutStyle)( style [**GoLayoutStyle**](api.GoLayoutStyle#) )\n" + 
"- [**SetMargin**](api.GoWindow#setMargin)( left [**int**](), top [**int**](), bottom [**int**](), right [**int**]() )\n" + 
"- [**SetOnClose**](api.GoWindow#setOnClose)( f [**func**]()() )\n" + 
"- [**SetOnConfig**](api.GoWindow#setOnConfig)( f [**func**]()() )\n" + 
"- [**SetPadding**](api.GoWindow#setPadding)( left [**int**](), top [**int**](), bottom [**int**](), right [**int**]() )\n" + 
"- [**SetPos**](api.GoWindow#setPos)( x [**int**](), y [**int**]() )\n" + 
"- [**SetSize**](api.GoWindow#setSize)( width [**int**](), height [**int**]() )\n" + 
"- [**SetSpacing**](api.GoWindow#setSpacing)( style [**GoLayoutSpacing**](api.GoLayoutSpacing#) )\n" + 
"- [**SetTitle**](api.GoWindow#setTitle)( title [**string**]() )\n" + 
"- [**Show**](api.GoWindow#show)()\n" + 
"- [**ShowModal**](api.GoWindow#showModal)()\n" + 
"- [**Size**](api.GoWindow#size)()  ( width [**int**](), height [**int**]() )\n" + 
"- [**Title**](api.GoWindow#title)()  ( title [**string**]() )\n" + 
"- [**Widget**](api.GoWindow#widget)()  ( widget [***GioWidget**](api.GioWidget#) )\n" + 
"\n---\n" + 

"###### **Detail Description** <a name=\"details\"></a>\n" + 

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
"\n\n" + 

"###### **GoWindowObj Properties** <a name=\"properties\"></a>\n\n" + 

"[type]() **GoWindowObj** [struct](){}\n" + 
"- [**GioObject**](api.GoObject#)\n" + 
"- [**GoSize**](api.GoSize#)\n" + 
"- [**GoPos**](#goPos)\n" + 
"- gio [***app_gio.Window**]()\n" + 
"- title [string]()\n" + 
"- frame [***GoLayoutObj**](api.GoLayout#)\n" + 
"- menubar [***GoMenuBarObj**](api.GoMenuBar#)\n" + 
"- statusbar [***GoLayoutObj**](api.GoLayout#)\n" + 
"- layout [***GoLayoutObj**](api.GoLayout#)\n" + 
"- eventmask [***GoEventMaskObj**](api.GoEventMask#)\n" + 
"- mainwindow [bool]()\n" + 
"- modalwindow [bool]()\n" + 
"- modalstyle [string]()\n" + 
"- ModalAction [int]()\n" + 
"- ModalInfo [string]()\n" + 
"- popupmenus **[[]*GoPopupMenuObj](api.GoPopupMenu#)**\n" + 
"- popupwindow [***GoPopupWindowObj**](api.GoPopupWindow#)\n" + 
"- onClose [func]()()\n" + 
"- onConfig [func]()()\n" + 
"\n\n" + 

"###### **Member Function Documentation**\n" + 

"<a name=\"goMainWindow\"></a> [**GoMainWindow**](api.GoWindow#goMainWindow)( windowTitle **string** )  ( hWin [***GoWindowObj**](#goWindowObj) )\n" + 
"- Create a new main window\n\n" + 
"<a name=\"goWindow\"></a> [**GoWindow**](api.GoWindow#goWindow)( windowTitle **string** )  ( hWin [***GoWindowObj**](#goWindowObj) )\n" + 
"- Create a new window\n\n" + 
"<a name=\"goModalWindow\"></a> [**GoModalWindow**](api.GoWindow#goModalWindow)( windowTitle **string** )  ( hWin [***GoWindowObj**](#goWindowObj) )\n" + 
"- Create a new modal window\n\n" + 
"<a name=\"addPopupMenu\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**AddPopupMenu(** popupMenu [***GoPopupMenuObj**](api.GoPopupMenuObj#) **)**\n" + 
"- Add a new popup menu.\n\n" + 
"<a name=\"centre\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**Centre()**\n" + 
"- Centre the window on the client screen.\n\n" + 
"<a name=\"clearPopupMenus\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**ClearPopupMenus()**\n" + 
"- Clear the popup menus.\n\n" + 
"<a name=\"clientPos\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**ClientPos(** x **int,** y **int )**\n" + 
"- Returns the client postion of the window.\nThis is usually set to (0,0)\n\n" + 
"<a name=\"clientSize\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**ClientSize() (** width **int,** height **int )**\n" + 
"- Returns the inner client size of the window.\n\n" + 
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
"<a name=\"menubar\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**MenuBar() (** menubar [***GoMenuBarObj**](api.GoMenuBar#) **)**\n" + 
"- Installs and returns a pointer to the window main menu bar.\n\n" + 
"<a name=\"menuPopup\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**MenuPopup() (** popupMenu [***GoPopupMenuObj**](api.GoPopupMenu#) **)**\n" + 
"- Returns a pointer to the popup menu at index idx.\n\n" + 
"<a name=\"minimize\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**Minimize()**\n" + 
"- Minimize the window.\n\n" + 
"<a name=\"modalStyle\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**ModalStyle() (** style **string )**\n" + 
"- Returns the modal style of a modal window.\n\n" + 
"<a name=\"objectType\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**ObjectType() (** type **string )**\n" + 
"- Returns the object type as a string definition \"GoWindowObj\".\n\n" + 
"<a name=\"popupWindow\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**PopupWindow() (** popupWindow [***GoPopupWindowObj**](api.GoPopupWindow#) **)**\n" + 
"- Returns a pointer to the windows popup window.\n\n" + 
"<a name=\"pos\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**Pos(** x **int, y int )**\n" + 
"- Returns the screen postion of the window.\n\n" + 
"<a name=\"raise\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**Raise()**\n" + 
"- Raises the window to be top most window.\n\n" + 
"<a name=\"refresh\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**Refresh()**\n" + 
"- Refresh the window.\n\n" + 
"<a name=\"setBorder\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**SetBorder(** style **[**GoBorderStyle**](api.GoBorderStyle#), width **int**, radius **int**, color [**GoColor**](api.GoColor#) **)**\n" + 
"- Add a border to the window.\n\n" + 
"<a name=\"setBorderColor\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**SetBorderColor(** color [**GoColor**](api.GoColor#) **)**\n" + 
"- Change the border color of the window border.\n\n" + 
"<a name=\"setBorderRadius\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**SetBorderRadius(** radius **int )**\n" + 
"- Change the border radius of the window border.\n\n" + 
"<a name=\"setBorderStyle\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**SetBorderStyle(** style [**GoBorderStyle**](api.GoBorderStyle#) **)**\n" + 
"- Change the border style of the window border.\n\n" + 
"<a name=\"setBorderWidth\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**SetBorderWidth(** width **int )**\n" + 
"- Change the border width of the window border.\n\n" + 
"<a name=\"setLayoutStyle\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**SetLayoutStyle(** style [**GoLayoutStyle**](api.GoLayoutStyle#) **)**\n" + 
"- Changes the window central layout style.\n\n" + 
"<a name=\"setMargin\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**SetMargin(** left **int,** top **int,** bottom **int,** right **int )**\n" + 
"- Sets the window margin sizes.\n\n" + 
"<a name=\"setOnClose\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**SetOnClose(** f **func() )**\n" + 
"- Adds a function to be called when the window is closed.\n\n" + 
"<a name=\"setOnConfig\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**SetOnConfig(** f **func() )**\n" + 
"- Adds a function to be called when the window is reconfigured.\n\n" + 
"<a name=\"setPadding\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**SetPadding(** left **int,** top **int,** bottom **int,** right **int )**\n" + 
"- Sets the window padding sizes.\n\n" + 
"<a name=\"setPos\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**SetPos(** x **int, y int )**\n" + 
"- Moves the position of the window on screen.\n\n" + 
"<a name=\"setSize\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**SetSize(** width **int, height int )**\n" + 
"- Resizes the window on screen.\n\n" + 
"<a name=\"setSpacing\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**SetSpacing(** style [**GoLayoutSpacing**](api.GoLayoutSpacing#) **)**\n" + 
"- Changes the window central layout widget spacing.\n\n" + 
"<a name=\"setTitle\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**SetTitle(** title **string )**\n" + 
"- Change the text displayed in the window title bar.\n\n" + 
"<a name=\"show\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**Show()**\n" + 
"- Activate the window loop and set as top window.\n\n" + 
"<a name=\"showModal\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**ShowModal()**\n" + 
"- Activate the window as a modal application window and set as topmost window.\n\n" + 
"<a name=\"size\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**Size() (** width **int,** height **int )**\n" + 
"- Returns the outer size of the window.\n\n" + 
"<a name=\"title\"></a> **(ob** [***GoWindowObj**](api.GoWindow#)**)**.**Title() (** title **string )**\n" + 
"- Return the text displayed by the window title bar.\n\n" + 
"<a name=\"widget\"></a> **(ob** [***GoWindowObj)**](api.GoWindow#).**Widget() (** widget [**GioWidget**](api.GioWidget#) **)**\n" + 
"- Returns a pointer to the window widget properties.\n\n"}

var GoButton []string = []string{"" + 
"The GoButton object provides an command button. [More...](#details)\n\n" + 
"[**GoButton**](api.GoButton#goButton)( parent [**GoObject**]( GoObject#), text [**string**]() ) ( hObj [***GoButtonObj**](api.GoButton#) )\n\n" + 
"member functions...\n\n" + 
"- [**Draw**](api.GoButton#draw)( gtx [**layout_gio.Context**]() )\n" + 
"- [**Layout**](api.GoButton#layout)( gtx [**layout_gio.Context**]() )\n" + 
"- [**ObjectType**](api.GoWindow#objectType)() ( type [**string**]() )\n" + 
"- [**SetFaceColor**](api.GoButton#setFaceColor) ( color [**GoColor**](api.GoColor#) )\n" + 
"- [**SetOnClick**](api.GoButton#setOnClick) ( f [**func**]() )\n" + 
"- [**SetText**](api.GoButton#setText) ( text [**string**]() )\n" + 
"- [**SetTextColor**](api.GoButton#setTextColor) ( color [**GoColor**](api.GoColor#) )\n" + 
"- [**Text**](api.GoButton#text)() ( text [**string**]() )\n" + 
"- [**Widget**](api.GoWindow#widget)() ( widget [***GioWidget**](api.GioWidget#) )\n\n" + 

"###### **Detail Description** <a name=\"details\"></a>\n" + 
"The GoButton object provides an command button.\n\n" + 

"###### **GoButtonObj Properties** <a name=\"properties\"></a>\n\n" + 

"[type]() **GoButtonObj** [struct](){}\n" + 
"- [**GioObject**](api.GoObject#)\n" + 
"- [**GioWidget**](api.GoWidget#)\n" + 
"- font [**font_gio.Font**]()\n" + 
"- fontSize [**unit_gio.Sp**]()\n" + 
"- text [string]()\n" + 
"- textColor [**GoColor**]()\n" + 
"- faceColor [**GoColor**]()\n" + 
"- cornerRadius [**unit_gio.Dp**]()\n" + 
"- inset [**layout_gio.Inset**]()\n" + 
"- shaper [***text_gio.Shaper**]()\n" + 
"- onClick [func]()()\n" + 
"\n\n" + 

"##### **Member Function Documentation**\n" + 

"<a name=\"goButton\"></a> **[GoButton](api.GoButton#) ( parent [*GoObject](api.GoObject#), text [string](string) ) ( hObj [*GoButtonObj](api.GoButton#) )**\n" + 
"- Add a new button object, assign a parent and caption text." + 
" Usually a button object will be added to a parent layout.\n\n" + 
"```\n" + 
"   example:   buttonExit := GoButton(parent, \"Exit\")\n" + 
"```\n" + 
"<a name=\"draw\"></a> **(ob** [***GoButtonObj)**](api.GoButton#).**Draw(gtx [layout_gio.Context](\"https://gioui.org/layout/context.go\")**\n" + 
"- Draw the button object. This function is usually called from a GoLayoutObj.Draw() function" + 
" during the GoWindowObj.render() event and not called directly by a user application. \n" + 
"It is necessary to declare it as a public function to allow new objects using the GoObject Interface to be created by the user application.\n\n" + 

"<a name=\"layout\"></a> **(ob** [***GoButtonObj)**](api.GoButton#).**Layout(gtx [layout_gio.Context](\"https://gioui.org/layout/context.go\") )**\n" + 
"- Layout the button object. This function is usually called from a GoButtonObj.Draw() function" + 
" during the GoWindowObj.render() event and not called directly by a user application. \n" + 
"It is necessary to declare it as a public function to facilitate drawing of new objects created by the user application.\n\n" + 

"<a name=\"objectType\"></a> **(ob [*GoButtonObj)](api.GoButton#).ObjectType() ( type [string](string) )**\n" + 
"- Returns the object type as a string definition \"GoButtonObj\".\n\n" + 

"<a name=\"setFaceColor\"></a> **(ob [*GoButtonObj)](api.GoButton#).SetFaceColor( color [GoColor](goColor) )**\n" + 
"- Set the button face color.\n\n" + 

"<a name=\"setOnClick\"></a> **(ob [*GoButtonObj)](api.GoButton#).SetOnClick( f [func](func) )**\n" + 
"- Set the button click function.\n\n" + 
"```\n" + 
"   example:   buttonExit.SetOnClick(ActionExit_Clicked)\n" + 
"```\n" + 
"<a name=\"setText\"></a> **(ob [*GoButtonObj)](api.GoButton#).SetText( text [string](string) )**\n" + 
"- Set the caption text of the button.\n\n" + 

"<a name=\"setTextColor\"></a> **(ob [*GoButtonObj)](api.GoButton#).SetTextColor( color [GoColor](goColor) )**\n" + 
"- Set the text color of the button caption.\n\n" + 

"<a name=\"text\"></a> **(ob [*GoButtonObj)](api.GoButton#).Text() ( text [string](string) )**\n" + 
"- Returns the button caption text.\n\n" + 

"<a name=\"widget\"></a> **(ob [*GoButtonObj)](api.GoButton#).Widget() ( widget [*GioWidget](api.GioWidget#) )**\n" + 
"- Returns a pointer to the button widget properties.\n\n" + 
"\n"}

var GoButtonGroup []string = []string{"" + 
"The GoButtonGroup object groups button objects. [More...](#details)\n\n" + 
"- [**GoButtonGroup**](api.GobuttonGroup#)() ( [***GoButtonGroupObj**](api.GoButtonGroup#properties) )\n\n" + 
"member functions...\n\n" + 
"- [**Value**](api.GoButtonGroup#value)() ([**string**]() )\n" + 

"###### **Detail Description** <a name=\"details\"></a>\n" + 

"###### **GoButtonGroupObj Properties** <a name=\"properties\"></a>\n\n" + 

"[type]() **GoButtonGroupObj** [struct](){}\n" + 
"- enum [***widget_int.GioEnum**]()\n" + 
"\n\n" + 

"##### **Member Function Documentation**\n" + 

"- Value() string " + 

"\n"}

var GoCanvas []string = []string{"" + 
"The GoCanvas object is a 2D view used to render GoCanvasItems. [More...](#details)\n\n" + 
"- [**GoCanvas**](api.GoCanvas#goCanvas)( parent [**GoObject**](api.GoObject#) )  ( hObj [***GoCanvasObj**].(api.GoCanvas#properties) )\n\n" +
"member functions...\n\n" + 
"###### **Detail Description** <a name=\"details\"></a>\n" + 

"###### **GoCanvasObj Properties** <a name=\"properties\"></a>\n\n" + 

"##### **Member Function Documentation**\n" + 

"\n"}

var GoCheckBox []string = []string{"" + 
"The GoCheckBox object provides a checkbox with a label. [More...](#details)\n\n" + 
"- [**GoCheckBox**](api.GoCheckBox#goCheckBox)( parent [**GoObject**](api.GoObject#), label [**string**]()) [***GoCheckBoxObj**](api.GoCheckBox#properties)\n\n" + 
"member functions...\n\n" + 

"###### **Detail Description** <a name=\"details\"></a>\n" + 



"[type]() **GoCheckBoxObj** [struct](){}\n" + 
"- [**GioObject**](api.GoObject#)\n" + 
"- [**GioWidget**](api.GoWidget#)\n" + 
"- checkable [**widget_int.GioCheckable**]()\n" + 
"- checkBox [***widget_gio.Bool**]()\n" + 

"##### **Member Function Documentation**\n" + 

"\n"}

var GoImage []string = []string{"" + 
"The GoImage object provides a view with a rendered image. [More...](#details)\n\n" + 
"- [**GoImage**](api.GoImage#goImage)( parent [**GoObject**](api.GoObject#), src [**string**]())  ( hObj [***GoImageObj**](api.GoImage#properties) )\n\n" + 
"member functions...\n\n" + 
"###### **GoImageObj Properties** <a name=\"properties\"></a>\n\n" + 

"[type]() **GoImageObj** [struct](){}\n" + 

"##### **Member Function Documentation**\n" + 

"\n"}

var GoLabel []string = []string{"" + 
"The GoLabel object provides a text label. [More...](#details)\n\n" + 
"- [**GoLabel**](api.GoLabel#goLabel)( parent [**GoObject**](api. GoObject#), text [**string**]())  ( hObj [***GoCheckBoxObj**](api.GoLabel#properties) )\n\n" + 
"member functions...\n\n" + 
"###### **GoLabelObj Properties** <a name=\"properties\"></a>\n\n" + 

"[type]() **GoLabelObj** [struct](){}\n" + 

"##### **Member Function Documentation**\n" + 

"\n"}

var GoLayout []string = []string{"" + 
"The GoLayout object is a container used to layout other objects. [More...](#details)\n\n" + 
"- [**GoLayout**](api.GoLayout#goLayout)( parent [**GoObject**](api.GoObject#), style [**GoLayoutStyle**](api.GoLayoutStyle#))  ( hObj [***GoLayoutObj**](api.GoLayout#properties) )\n" + 
"- [**GoBoxLayout**](api.GoLayout#goBoxLayout)( parent [**GoObject**](api.GoObject#), style [**GoLayoutStyle**](api.GoLayoutStyle#))  ( hObj [***GoLayoutObj**](api.GoLayout#properties) )\n" + 
"- [**GoHBoxLayout**](api.GoLayout#goHBoxLayout)( parent [**GoObject**](api.GoObject#) )  ( hObj [***GoLayoutObj**](api.GoLayout#properties) )\n" + 
"- [**GoVBoxLayout**](api.GoLayout#goVBoxLayout)( parent [**GoObject**](api.GoObject#) )  ( hObj [***GoLayoutObj**](api.GoLayout#properties) )\n" + 
"- [**GoFlexBoxLayout**](api.GoLayout#goFlexBoxLayout)( parent [**GoObject**](api.GoObject#), style [**GoLayoutStyle**](api.GoLayoutStyle#) )  ( hObj [***GoLayoutObj**](api.GoLayout#properties) )\n" + 
"- [**GoHFlexBoxLayout**](api.GoLayout#goHFlexBoxLayout)( parent [**GoObject**](api.GoObject#) )  ( hObj [***GoLayoutObj**](api.GoLayout#properties) )\n" + 
"- [**GoVFlexBoxLayout**](api.GoLayout#goVFlexBoxLayout)( parent [**GoObject**](api.GoObject#) )  ( hObj [***GoLayoutObj**](api.GoLayout#properties) )\n" + 
"- [**GoPopupMenuLayout**](api.GoLayout#goPopupMenuLayout)( parent [**GoObject**](api.GoObject#) )  ( hObj [***GoLayoutObj**](api.GoLayout#properties) )\n\n" + 

"###### **GoLayoutObj Properties** <a name=\"properties\"></a>\n\n" + 

"[type]() **GoLayoutObj** [struct](){}\n" + 
"- [**GioObject**](api.GoObject#)\n" + 
"- [**GioWidget**](api.GoWidget#)\n" + 
"- Scrollbar [**GoScrollbar**](api.GoScrollBar#)\n" + 
"- AnchorStrategy\n" + 
"- list_gio [***widget_gio.List**]()\n" + 
"- flex_gio [***layout_gio.Flex**]()\n" + 
"- style [**GoLayoutStyle**](api.GoLayoutStyle#)\n" + 
"- flexControls **[[]layout_gio.FlexChild]()**\n" + 

"##### **Member Function Documentation**\n" + 

"\n"}

var GoList []string = []string{"" + 
"- [**GoList**](api.GoList#goList)( parent [**GoObject**](api.GoObject#) )  ( hObj [***GoListViewItemObj**](api.GoListViewItem#properties))\n" + 


"\n"}

var GoListView []string = []string{"" + 
"The GoListView object implements a list or tree view. [More...](#details)\n\n" + 
"- [**GoListView**](api.GoListView#goListView)( parent [**GoObject**](api.GoObject#) )  ( hObj [***GoListViewObj**](api.GoListView#properties))\n" + 

"###### **GoListViewObj Properties** <a name=\"properties\"></a>\n\n" + 

"[type]() **GoListViewObj** [struct](){}\n" + 

"- [**GioObject**](api.GoObject#)\n" + 
"- [**GioWidget**](api.GoWidget#)\n" + 
"- Scrollbar [**GoScrollbar**](api.GoScrollBar#)\n" + 
"- AnchorStrategy\n" + 
"- state [***widget_gio.List**]()\n" + 
"- columns [int]()\n" + 
"- itemColor [**GoColor**]()\n" + 
"- itemSize [int]()\n" + 
"- currentItem [***GoListViewItemObj**]()\n" + 
"- layout [***GoLayoutObj**]()\n" + 
"- onItemClicked [**func**]()( [[]int]() )\n" + 
"- onItemDoubleClicked [**func**]()( [[]int]() )\n" + 

"\n"}

var GoListViewItem []string = []string{"" + 
"The GoListViewItem object implements an item in a list view. [More...](#details)\n\n" + 
"- [**GoListViewItem**](api.GoListViewItem#goListViewItem)( parent [**GoObject**](api.GoObject#), data **[[]byte]()**, text [**string**](), listLevel [**int**](), listId [**int**]() )  ( hObj [***GoListViewItemObj**](api.GoListViewItem#properties))\n" + 

"###### **GoListViewItemObj Properties** <a name=\"properties\"></a>\n\n" + 

"[type]() **GoListViewItemObj** [struct](){}\n" + 

"- [**GioObject**](api.GoObject#)\n" + 
"- [**GioWidget**](api.GoWidget#)\n" + 
"- color [**GoColor**](api.GoColor#)\n" +   // foreground color
"- expanded [bool]()\n" +       // true if the tree node is expanded
"- font [font_gio.Font]()\n" +  // label font
"- fontSize [unit_gio.Sp]()\n" + // label font size
"- icon [[]byte]()\n" +          // icon svg data
"- iconColor [**GoColor**](api.GoColor#)\n" +                // icon color
"- iconSize [int]()\n" +                     // size of icon determines height of listViewItem
"- id [int]()\n" +                           // position of listViewItem in parentItem
"- label [string]()\n" +                     // text displayed
"- level [int]()\n" +                        // tree level 0 ...
"- listView [***GoListViewObj**](api.GoListView#)\n" +          // view to display all listViewItems
"- maxlines [int]()\n" + 
"- shaper [***text_gio.Shaper**]()\n\n" + 
"// Cached values.\n" + 
"- op [**paint_gio.ImageOp**]()\n" + 
"- imgSize  [int]()\n" + 
"- imgColor [**GoColor**](api.GoColor#)\n" + 

"\n"}

var GoMenu []string = []string{"" + 
"The GoMenu object provides a new menu item to be inserted into a menubar. [More...](#details)\n\n" + 
"\n"}

var GoMenuBar []string = []string{"" + 
"The GoMenuBar object provides a bar into which main menu labels are inserted. [More...](#details)\n\n" + 
"\n"}

var GoMenuItem []string = []string{"" + 
"The GoMenuItem object is a menu label to which an action is associated. [More...](#details)\n\n" + 
"\n"}

var GoRadioButton []string = []string{"" + 
"The GoRadioButton object provides a radio button with a label. [More...](#details)\n\n" + 
"\n"}

var GoRichText []string = []string{"" + 
"The GoRichText object provides a richtext display object. [More...](#details)\n" + 
"\n"}

var GoSlider []string = []string{"" + 
"The GoSlider object provides a sliding bar selection control. [More...](#details)\n" + 
"\n"}

var GoSpacer []string = []string{"" + 
"The GoSpacer object is used to introduce padding between controls in a GoLayout. [More...](#details)\n" + 
"\n"}

var GioObject []string = []string{"" + 
"The GioObject is the base object embedded in all utopia controls. [More...](#details)\n\n" + 

"[**GioObject**](#properties){ parent [**GoObject**](api.GoObject#), parentWindow [**GoWindowObj**](api.GoWindow#), []GoObject **[[]GoObject](api.GoObject#)**, sizePolicy [**GoSizePolicy**](api.GoSizePolicy#) }\n\n" + 
"member functions...\n\n" + 
"- [**AddControl**](api.GioObject#addControl)( control [**GoObject**](api.GoObject#) )\n" + 
"- [**Clear**](api.GioObject#clear)()\n" + 
"- [**DeleteControl**](api.GioObject#deleteControl)( control [**GoObject**](api.GoObject#properties) )\n" + 
"- [**InsertControl**](api.GioObject#insertControl)( control [**GoObject**](api.GoObject#properties), idx [int]() )\n" + 
"- [**Objects**](api.GioObject#objects)()  ( objects [**[]GoObject](api.GoObject#properties)** )\n" + 
"- [**ParentControl**](api.GioObject#parentControl)()  ( control [**GoObject**](api.GoObject#properties) )\n" + 
"- [**ParentWindow**](api.GioObject#parentWindow)()  ( window [**GoWindowObj**](api.GoWindow#properties) )\n" + 
"- [**RemoveControl**](api.GioObject#removeControl)( control [**GoObject**](api.GoObject#properties) )\n" + 
"- [**SetConstraints**](api.GioObject.setConstraints)( size [**GoSize**](api.GoSize#), cs [layout_gio.Constraints]() ) ( cs [layout_gio.Constraints]() )\n" + 
"- [**SizePolicy**](api.GioObject#sizePolicy)()  ( sizePolicy [***GoSizePolicy**](api.GoSizePolicy#) )\n" + 
"- [**SetHorizSizePolicy**](api.GioObject#setHorizSizePolicy)( horiz [**GoSizeType**](api.GoSizeType#) )\n" + 
"- [**SetSizePolicy**](api.GioObject#setSizePolicy)( horiz [**GoSizeType**](api.GoSizeType#), vert [**GoSizeType**](api.GoSizeType#) )\n" + 
"- [**SetVertSizePolicy**](api.GioObject#setVertSizePolicy)( vert [**GoSizeType**](api.GoSizeType#) )\n" + 
"- [**Widget**](api.GoObject#widget)()  ( widget [**GioWidget**](api.GioWidget#) )\n" + 


"###### **Detail Description** <a name=\"details\"></a>\n" + 

"The GioObject is the embedded object in all utopia controls. Inheritance of Utopia controls is performed through the GioObject.\n\n" + 
"All GioObjects receive a parent GoObject reference and a parent GoWindowObj reference in their constructor.\n\n" + 
"When a control is created with an object as parent, the parents GioObject will automatically insert the new control object into it's Objects() list.\n\n" + 
"Controls can be added, inserted or removed from the GioObject producing object trees. The controls are then recursively layed out by traversing these object trees, usually performed by the GoLayout controls, to layout all of their contents.\n\n" + 



"###### **GioObject Properties** <a name=\"properties\"></a>\n\n" + 

"[**GioObject**](#properties){ **Parent**: parent [**GoObject**](api.GoObject#), **Window**: parentWindow [**GoWindowObj**](api.GoWindow#), **Controls**: **[[]GoObject](api.GoObject#)**, **GoSizePolicy**: [**sizePolicy**](api.GoSizePolicy#) }\n\n" + 
"[type]() **GioObject** [struct](){}\n" + 

"- Parent [**GoObject**](api.GoObject#)\n" + 
"- Window [***GoWindowObj**](api.GoWindow#)\n" + 
"- Controls **[[]GoObject](api.GoObject#)**\n" + 
"- GoSizePolicy [***GoSizePolicy**](api.GoSizePolicy#)\n" + 


"###### **Member Function Documentation**\n" + 

"<a name=\"addControl\"></a> **(ob** [**GioObject**](api.GioObject#)**)**.**AddControl(** control [**GoObject**](api.GoObject#) **)**\n" + 
"- Add a new control to parent.\n\n" + 
"<a name=\"clear\"></a> **(ob** [**GioObject**](api.GioObject#)**)**.**Refresh()**\n" + 
"- Remove and delete all the controls from the parent control.\n\n" + 
"<a name=\"deleteControl\"></a> **(ob** [**GioObject**](api.GioObject#)**)**.**DeleteControl(** control [**GoObject**](api.GoObject#) **)**\n" + 
"- Remove and delete a control from the parent control.\n\n" + 
"<a name=\"insertControl\"></a> **(ob** [**GioObject**](api.GioObject#)**)**.**InsertControl(** control [**GoObject**](api.GoObject#), index [**int**]() **)**\n" + 
"- Insert a new control in parent at position index.\n\n" + 
"<a name=\"objects\"></a> **(ob** [**GioObject)**](api.GioObject#).**Objects() (** controls **[[]GoObject](api.GoObject#)** **)**\n" + 
"- Returns a pointer to the window widget properties.\n\n" + 
"<a name=\"parentControl\"></a> **(ob** [**GioObject)**](api.GioObject#).**ParentControl() (** parent [**GoObject**](api.GoObject#) **)**\n" + 
"- Returns a pointer to the controls parent control.\n\n" + 
"<a name=\"parentWindow\"></a> **(ob** [**GioObject)**](api.GioObject#).**ParentWindow() (** parent [**GoWindow**](api.GoWindow#) **)**\n" + 
"- Returns a pointer to the controls parent window.\n\n" + 
"<a name=\"removeControl\"></a> **(ob** [**GioObject**](api.GioObject#)**)**.**RemoveControl(** control [**GoObject**](api.GoObject#) **)**\n" + 
"- Remove but not delete a control from the parent control.\n\n" + 
"<a name=\"setConstraints\"></a> **(ob** [**GioObject**](api.GioObject#)**)**.**SetConstraints(** size [**GoSize**](api.GoSize#), constraints [**layout_gio.Constraints**]() **)  (** constraints [**layout_gio.Constraints**]() **)**\n" + 
"- Apply the controls max and min sizes adjusting for the controls size policy.\n\n" + 
"<a name=\"sizePolicy\"></a> **(ob** [***GioObject**](api.GioObject#)**)**.**SizePolicy() (** sizePolicy [***GoSizePolicy**](api.GoSizePolicy#) **)**\n" + 
"- Returns the size policy of the control.\n\n" + 
"<a name=\"setHorizSizePolicy\"></a> **(ob** [**GioObject**](api.GioObject#)**)**.**SetHorizSizePolicy(** horiz [**GoSizeType**](api.GoSizeType#) **)**\n" + 
"- Set the controls size policy in the horizontal direction. The vertical size policy remains unaltered.\n\n" + 
"<a name=\"setVertSizePolicy\"></a> **(ob** [**GioObject**](api.GioObject#)**)**.**SetVertSizePolicy(** vert [**GoSizeType**](api.GoSizeType#) **)**\n" + 
"- Set the controls size policy in the vertical direction. The horizontal size policy remains unaltered.\n\n" + 
"<a name=\"setSizePolicy\"></a> **(ob** [**GioObject**](api.GioObject#)**)**.**SetSizePolicy(** horiz [**GoSizeType**](api.GoSizeType#), vert [**GoSizeType**](api.GoSizeType#) **)**\n" + 
"- Set the controls size policy in both the horizontal and vertical directions.\n\n" + 
"\n"}


var GioWidget []string = []string{"" + 
"The GioWidget is the base widget object embedded in all utopia objects. [More...](#details)\n" + 
"- **[ChangeFocus](api.GioWidget#changeFocus)() (** focus **bool )**\n" + 
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
"- **[SetHeight](api.GioWidget#setHeight)(** height **int )**\n" + 
"- **[SetWidth](GioWidget#setWidth)(** width **int )**\n" + 

"##### **Detail Description** <a name=\"details\"></a>\n" + 

"The GioWidget represents the controls user interface. It receives mouse, keyboard and other events from the operating system, and determines how it is represented on the screen. " + 
" A GioWidget is clipped by its parent, usually a layout object, which is also a Giowidget in its own right.\n\n" + 

"If you wish to use an object to display child objects you will probably want to add a layout object to the parent object. (See Layouts.)\n\n" + 

"The GioWidget structure is embedded by other user controls to allow them to access all the GioWidget properties.\n\n" + 
"Importantly mouse and gesture events are activated by the underlying UtopiaGio framework, as and when the user control allocates a function to the GioWidget's event.\n" + 
"Key press and release events are activated by the underlying UtopiaGio framework, when the user control sets key filters and allocates a function to either of the widgets key press or release event.\n\n" + 

"GioWidget inherits GoMargin, GoBorder and GoPadding along with GoSize. Other properties are for visibility and assigning colour attributes for background, foreground, highlighting etc.\n" + 
" See the section on the GoColor class.\n\n" + 

"It also connects the Utopia framework to the Gio backend. All events are directed through the GioWidget class to each individual object, by declaring which events are to be monitored.\n" + 
" This is achieved by assigning a new procedure to the GioWidget using the SetOnPointerClick(), SetOnPointerPress(), SetOnPointerRelease(), SetOnPointerEnter(), SetOnPointerLeave() etc.\n" + 
" Key Events likewise can be activated with SetOnKeyPress(), SetOnKeyRelease(), and SetOnKeyEvent().\n\n" + 

"###### **GioWidget Properties** <a name=\"properties\"></a>\n\n" + 

"[type]() **GioWidget** [struct](){}\n" + 

""}