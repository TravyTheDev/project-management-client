# Project Management

Users can login to a desktop application, create a project and sub-tasks, assign users, edit, and leave private notes for themselves.

Users get notifications in real time about assignment and can live edit with other desktop users.

Manage projects on a gantt chart.

https://traviskleckley.com/project_management/
![](https://raw.githubusercontent.com/TravyTheDev/personal-site/refs/heads/main/public/images/live-edit.gif)

This is a desktop application made with Wails, Go and TypeScript(Vue).
It connects to a Go server https://github.com/TravyTheDev/personal-site-server which has a Sqlite database.

This started because I was wondering if it was possible to do access/refresh token auth on the desktop and it just got out of hand.

### Usage

You must have an account the admin has made for you to log in. 

When live editing with someone the inputs are debounced 300ms.

When focusing an input, it will save on blur. 

Sub-tasks will delete when the parent task is deleted. 

Users are notified as soon as the assignee input is filled.

### Optimizations 

Currently the data is not paginated.

I want to rewrite the gantt chart from scratch and not use a library mostly because I want floating labels on scroll.

When you logout it does a janky reboot because I'm using a library for the server sent events stream so I want to rewrite it myself so I can easily unsubscribe.

i18n.

### Lessons learned

Server to server websockets. 

Golang cookie jar and server to server requests.

You really are spoiled using an actual browser with it's nice APIs. 
