# Groupie Trackers

## **Overview**
Groupie Trackers is a web application that integrates with a given API to display detailed information about artists and bands. It features an interactive, user-friendly interface and real-time client-server communication for dynamic data presentation.

---

## **Features**
- **Artist Information**: View details such as name, image, members, active years, and first album.
- **Concert Locations & Dates**: Explore past and upcoming concert details.
- **Dynamic Event Handling**: Trigger server requests via user actions (e.g., search, filters).
- **Responsive Design**: Fully functional across various devices.

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
- **Frontend**: HTML, CSS (optional: JavaScript for interactivity)
- **API Format**: JSON

---

## **Setup Instructions**
1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   cd <repository-folder>
   ```

	2.	Run the Server:

```
go run .
```

	3.	Access the Application:
	•	Open your browser and navigate to http://localhost:8080.

Usage
	1.	Browse Artist Information:
	•	View artist details on the homepage.
	2.	Explore Concerts:
	•	Check out concert locations and dates via interactive UI elements.
	3.	Search/Filter:
	•	Use the search bar or dropdowns to dynamically filter data.
	4.	Event Handling:
	•	Trigger client-server communication by interacting with various UI elements.

Project Requirements
	•	Backend: Written entirely in Go.
	•	Frontend: Responsive and user-friendly website.
	•	Error Handling: Robust and resilient against crashes.
	•	Testing: Includes unit tests for key functionalities.
	•	Good Practices: Clean, maintainable, and well-documented code.

Allowed Go Packages
	•	io
	•	log
	•	os
	•	fmt
	•	net
	•	errors
	•	strings

Learning Outcomes

This project enhances skills in:
	•	Integrating and manipulating APIs.
	•	Designing client-server architectures in Go.
	•	Creating dynamic and responsive web applications.
	•	Writing clean and testable code.




