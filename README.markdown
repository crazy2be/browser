Browser Package
===============

What is it?
-----------
This package provides specific information about a user's browser, grabbed from the user-agent string. It can tell you what platform a user is using (mobile, desktop), and what browser they are using (firefox, IE).

Getting Started
---------------
First, install it:

    goinstall github.com/crazy2be/browser

Then, import it:

    import "github.com/crazy2be/browser"

Finally, use it:

    if browser.IsMobile(r) {
        ...
    } else {
       ...
    }