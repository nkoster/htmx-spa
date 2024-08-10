## HTMX Single Page Application

This repository hosts a Go-based server application that dynamically serves HTML content based on URL
paths using the `htmx` library to enhance interactivity without requiring full page reloads.
The application supports both frontend and backend routing, with full history support to ensure
seamless user navigation and state management.

### Project Structure

- `ui/`: Contains the `index.html` file that serves as the base layout for all pages.
- `templates/`: Holds HTML files (`home.html`, `page_a.html`, `page_b.html`) that are dynamically loaded into the base layout based on the URL.
- `static/`: Directory for static files such as images, CSS, and JavaScript files.

### Features

- **Dynamic Content Loading**: Dynamically inserts content from the `templates/` directory into the `index.html` layout depending on the URL path accessed.
- **HTMX Integration**: Utilizes `htmx` for dynamic interactions, allowing for partial page updates and enhancing user experience with minimal overhead.
- **Static File Serving**: Efficiently manages and serves static resources from the `static/` directory under the `/static/` URL path.

### How It Works

1. **Base HTML Serving**:
   - The root URL (`/`) serves the `index.html` file along with the `home.html` content integrated, providing the initial page layout and content.

2. **Dynamic Template Insertion**:
   - `/page_a.html` loads `index.html` with the content from `page_a.html`.
   - `/page_b.html` loads `index.html` with the content from `page_b.html`.

3. **Static Content**:
   - Accessible under `/static/`, allowing images, CSS, and JS to be loaded as needed.

### Setup and Running

#### Prerequisites

- Ensure you have Go installed on your system.

#### Running the Server

1. Clone the repository to your local machine.
2. Navigate to the cloned directory.
3. Run the server using:

   ```bash
   go run main.go

4. Access the application in a web browser at http://localhost:8080.

### Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes or enhancements.

### License

This project is licensed under the MIT License.
