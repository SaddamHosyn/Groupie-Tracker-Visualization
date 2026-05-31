# Groupie Trackers

## **Overview**

Groupie Trackers is a web application that integrates with a given API to display detailed information about artists and bands. It features an interactive, user-friendly interface and real-time client-server communication for dynamic data presentation.

---

## **Features**

- **Artist Information**: View details such as name, image, members, active years, and first album.
- **Concert Locations & Dates**: Explore past and upcoming concert details.
- **Real-time Search & Filter**: Dynamically search and filter artists by name with instant results.
- **Dynamic Event Handling**: Trigger server requests via user actions (e.g., search, filters).
- **Responsive Design**: Fully functional and beautifully styled across various devices with Tailwind CSS.

---

## **API Details**

The project uses an API with the following endpoints:

1. **Artists**: Basic information about artists/bands.
2. **Locations**: Concert location details.
3. **Dates**: Concert dates (past/upcoming).
4. **Relations**: Links the above data for seamless integration.

---

## **Tech Stack**

- **Language**: Go (backend)
- **Frontend**: HTML5, Tailwind CSS, JavaScript
- **Styling**: Tailwind CSS (utility-first CSS framework)
- **API Format**: JSON
- **Containerization**: Docker (multi-stage build)

---

## **Setup Instructions**

1. **Clone the Repository**:

   ```bash
   git clone <repository-url>
   cd <repository-folder>
   ```

2. **Run the Server**:

   ```bash
   go run .
   ```

3. **Access the Application**:
   - Open your browser and navigate to `http://localhost:8080`
   - The application will load with all artists displayed
   - Use the search bar to find specific artists

---

## **Docker Setup**

### **Option 1: Using Docker Compose (Recommended)**

The easiest way to run the application:

```bash
docker-compose up -d
```

To stop the application:

```bash
docker-compose down
```

To view logs:

```bash
docker-compose logs -f
```

### **Option 2: Manual Docker Build & Run**

**Build the Docker Image**:

```bash
docker build -t groupie-tracker:latest .
```

**Run the Container**:

```bash
docker run -p 8080:8080 groupie-tracker:latest
```

Then access the application at `http://localhost:8080`

### **Docker Architecture**

This project uses a **multi-stage Docker build** for optimal image size:

1. **Stage 1 (Builder)**:
   - Uses `golang:1.21-alpine`
   - Builds the Go binary
   - Downloads dependencies
   - Compiles the application

2. **Stage 2 (Runtime)**:
   - Uses minimal `alpine:latest` image
   - Copies only the binary and static files
   - Runs the compiled application
   - Significantly smaller final image size

---

1. **Browse Artist Information**:
   - View artist cards on the homepage with image, name, creation date, and first album.
2. **Search Artists**:
   - Use the search bar to filter artists by name in real-time.
   - Search results update instantly as you type.
   - See "No artists found" message if no matches are found.
3. **Explore Artist Details**:
   - Click on any artist card to view comprehensive details.
   - View all band members as styled badges.
   - Explore concert locations and tour dates.
   - Check previous performance history.
4. **Responsive Navigation**:
   - Sticky back button for easy navigation.
   - Smooth transitions and hover effects for better UX.

## **Project Requirements**

- **Backend**: Written entirely in Go.
- **Frontend**: Responsive and user-friendly website built with Tailwind CSS.
- **Search Functionality**: Real-time artist search and filtering.
- **UI/UX**: Modern, gradient-styled cards with smooth animations and transitions.
- **Error Handling**: Robust and resilient against crashes.
- **Testing**: Includes unit tests for key functionalities.
- **Good Practices**: Clean, maintainable, and well-documented code.

Allowed Go Packages
• io
• log
• os
• fmt
• net
• errors
• strings

---

## **UI/Design Features**

- **Modern Styling**: Built with Tailwind CSS for a clean, professional appearance.
- **Gradient Colors**: Yellow and purple gradient themes throughout the application.
- **Responsive Grid**: Adaptive layouts that adjust from 1 column on mobile to 4 columns on desktop.
- **Interactive Elements**: Smooth hover effects, scale animations, and shadow transitions.
- **Accessibility**: Semantic HTML and ARIA labels for better accessibility.
- **Dark Theme**: Modern dark background with high contrast text for better readability.

---Learning Outcomes

This project enhances skills in:
• Integrating and manipulating APIs.
• Designing client-server architectures in Go.
• Creating dynamic and responsive web applications.
• Writing clean and testable code.
