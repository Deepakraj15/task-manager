# Task Manager

## Description

**Task Manager** is a high-performance scheduling system built in Go.  
It allows you to schedule tasks to run at specific times using a **hierarchical time wheel** for efficient timer management.  
The scheduler supports seconds, minutes, and hours wheels, enabling tasks to be executed with minimal delay and high scalability.

NOTE: I did this project as a medium of learning GO Lang.
---

## Features

- Schedule tasks via HTTP API.
- Hierarchical time wheel implementation (seconds → minutes → hours).
- Supports delayed execution of tasks.
- Thread-safe and efficient for thousands of concurrent tasks.

---

## Tech Stack

- **Backend**: Go
- **Routing**: Chi (`github.com/go-chi/chi/v5`)
- **Logging**: Custom logger
- **Data Structures**: Hierarchical time wheel (slots & tasks)
- **Time Handling**: `time` package for scheduling

---

## Installation

1. Clone the repository:

```bash
git clone https://github.com/Deepakraj15/task-manager.git
cd task-manager