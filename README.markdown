Browser Package
===============

What is it?
-----------
This package provides specific information about a user's browser, grabbed from the user-agent string. It can tell you what platform a user is using (mobile, desktop), and what browser they are using (firefox, IE).

Getting Started
---------------

    goinstall github.com/crazy2be/browser


    import "github.com/crazy2be/browser"


    if browser.IsMobile(r) {
        ...
    } else {
       ...
    }