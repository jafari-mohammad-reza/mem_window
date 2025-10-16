# Mem Window

Mem Window is a **Go library** for managing memory in discrete "windows" with fine-grained control. It allows developers to allocate, monitor, and manipulate memory segments based on **timestamp, size, or custom conditions**. The library supports compression, custom code execution on window triggers, and both fixed and dynamic memory allocation.

---

## Features

- **Custom Memory Windows:** Allocate memory based on a fixed number of windows or let it grow dynamically as needed.
- **Conditional Execution:** Run custom functions when a window meets a specific condition (e.g., size threshold, timestamp, or custom logic).
- **Compression Support:** Reduce memory footprint by compressing windowed data efficiently.
- **Flexible Window Management:** Manage memory slices with precision, improving performance for data streaming, buffering, or analytics tasks.
- **Extensible Architecture:** Integrate custom logic for specific use cases without modifying core library code.

---

## Installation

Make sure you have [Go](https://golang.org/dl/) installed (version 1.20+ recommended).

##Create a Memory Window Manager

// Example: Initialize manager with 5 fixed windows
manager := mem_window.NewManager(5)

// Add data to a window
manager.AddData("window1", []byte("sample data"))

// Execute custom function on condition
manager.OnCondition(func(window mem_window.Window) {
    // custom logic
}, mem_window.ConditionSize(1024)) // trigger when window size >= 1024 bytes

Project Structure

internal/: Core library logic, not exposed outside the module.

pkg/: Public API for interacting with memory windows.

example/: Sample code demonstrating typical usage scenarios.

README.md: Documentation for setup and usage.

go.mod: Go module file with dependencies.

Contributing

Contributions are welcome! You can:

Improve functionality or performance.

Add new features for window management or compression.

Expand examples and documentation.

Please fork the repository and submit a Pull Request.

License

This project is licensed under the MIT License. See the LICENSE
 file for details.

Notes

Designed for Go developers needing precise memory window management.

Ideal for streaming data, analytics pipelines, buffering, or event-driven systems.

Supports both fixed-size and dynamic memory allocation strategies.
