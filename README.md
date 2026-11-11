# Introduction
So today I have gotten into GUI in go somehow, and I rembemer from my 1st few days of programming in python I have always wanted to make to-do list. So that is exactly the thing I am gonna do now. I know this project is kinda simple really I have been kinda dumb before.

# Plan
So my plan to achieve this is pretty simple, I will make GUI in library I found the simplies the process of creating GUI in go called `fyne` and I will use this library to make stuff like buttons for adding tasks, buttons for removing tasks, marking tasks as done and so on. The logic behind this is pretty simple and it shouldn't be that compliicated, the only part I am kinda worried about is making the mark as compleate button for every task, displaying the task and reading tasks from a file to persist them.

# Details
I will be able to use AI, but not just vibe code and done, I want it to be mine so I will code it myself and I will use AI for debugging or generally for help with syntax.
Once again I will specify what features will this app have:
1. Buttons for adding and removing tasks.
2. Text field for user to type to.
3. Mark button to mark as compleate or not (1/0)
4. Persistent storage in a file, the file will be updated after every action e.g. adding a task.
5. A theme (purple/black)

As I have never done a GUI repository or a bigger project than simple tkinter windows the developmnet of this will surely be slow as fuck. I will try my best to do it in few days or so but it will depend on a lot of factors how quickly this will be done. Also by the time I am creating this repository I have not finished `search_engine` ranking system which I am really looking forward to so I will have to do that too while making this project.

# Updates
Now this project is in healthy state, I have not finished the marking and persistent functions, but everything else is working fine, I have also not even thought about how I want to make the marking system since I it is pretty weird to work with, or at least by how I think about it. I have thought about making it like splitted into 2 `container.NewVBox()` 1 will hold the name of task and 2nd will hold 1 or 0 depending on whether it is marked as done or not but I think it might not reflect the list well I could make something different like making the task contain `[x]` using `string.Join` or how its called if it works like this since I have not used this any time, so I think it might be good to use it to learn it since it could be used in more project for easier logic fruther on.
