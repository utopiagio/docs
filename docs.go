// SPDX-License-Identifier: Unlicense OR MIT

/* github.com/utopiagio/docs/docs.go */

package docs

import (
	"github.com/utopiagio/docs/apireference"
	"github.com/utopiagio/docs/reference"
)


var Overview []string = []string{"[**UtopiaGio**](https://github.com/utopiagio/utopia) is a Go framework library module calling the [**Gio**](https://gioui.org) library module. Gio is a cross-platform immediate mode GUI.\n\n" + 
"The [**GoApplication**](GoApplication#) class/structure maintains a list of GoWindows and manages the control of the GoWindows and their running threads.\n\n" +
"Each [**GoWindow**](GoWindow#) runs it's own message loop, but it will be possible to send and receive communications over channels between windows.\n\n" +
"The framework allows the building of more complex programs without the necessity to write code calling the Gio module directely. In turn this means reduced calls to Gio directly, but with the ability to still write specific Gio routines. It is also possible to use all of the Gio widgets by encapsulating within a new UtopiaGio widget structure and calling the Layout function from the GoObject.Draw() function.\n\n" +
"Inheritance is achieved by embedding the new [**GioObject**](GioObject#), and access to the user interface is provided by embedding the new [**GioWidget**](GioWidget#)." +
"New layout methods have been introduced requiring a very small change to the Gio layout package. The Gio widget package is still used with some of the widgets, but the intention is to move any relevant code for GoWidgets to the **utopia/internal/widget** package.\n\n" +
"Access to the underlying OS Screen and Main Window is provided through the **utopia/desktop** package, wwith the possibily to retrieve position, size and scaling of gio windows. The [**Pos**](GoWindow#pos) function has been added to the Gio package, which along with the [**Size**](GoWindow#size) function allows positioning and sizing of the gio window. Also available at run time using the utopia.GoWindowObj [**SetPos()**](GoWindow#setPos) and [**SetSize()**](GoWindow#setSize) functions.\n\n"}


func Page(packg string, doc string) (content []string) {
	switch packg {
		case "api":
			return apireference.Page(doc)
		case "ref":
			return reference.Page(doc)
		default:
			return Overview
	}
}
