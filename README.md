# Hotels Data Merge

[![Go Version](https://img.shields.io/badge/Go-1.24.2-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

[![Test Cases](https://github.com/robinloh/hotelsDataMerge/workflows/Test%20Cases/badge.svg)](https://github.com/robin/hotelsDataMerge/actions/workflows/testing.yml)
[![golangci-lint](https://github.com/robinloh/hotelsDataMerge/workflows/golangci-lint/badge.svg)](https://github.com/robin/hotelsDataMerge/actions/workflows/golangci-lint.yml)


## Summary

A backend application to aggregate, clean, and merge hotel data from multiple external suppliers (Acme, Patagonia, and Paperflies). 

The application fetches raw hotel data from various supplier APIs every 5 seconds, parses and normalizes the data into a unified format, merges duplicate hotels, and provides a clean, consolidated data through both gRPC and REST API endpoints.

## 2. Stack

**Backend Language:** 
Go 1.24

**Architecture:** Microservice with gRPC + REST Gateway


## 3. Installation Prerequisites

**System Requirements:**
- Go 1.24 or higher
- Internet connection (for fetching supplier data)

**Go Installation:**
```bash
# macOS (using Homebrew)
brew install go

# Linux
wget https://golang.org/dl/go1.24.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.24.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# Windows
# Download from https://golang.org/dl/ and follow installation instructions
```

**Verify Installation:**
```bash
go version
# Should output: go version go1.24.x darwin/amd64 (or similar)
```

## 4. How to Run this Application

**1. Clone the Repository:**
```bash
git clone <repository-url>
cd hotelsDataMerge
```

**2. Install Dependencies:**
```bash
go mod download
```

**3. Run the Application:**
```bash
go run main.go
```

**4. Verify Services are Running:**
- gRPC Server: `localhost:8080`
- REST API Gateway: `localhost:8090`

**5. Check Logs:**
The application will output structured logs showing:
- Service startup information
- Supplier data fetching status
- Data processing progress
- Any errors or warnings

**Note:** 
The application automatically starts fetching supplier data every 5 seconds in the background.

## 5. APIs

### 5.1. Table of APIs

| Endpoint | Method | Protocol | Description | Request Parameters | Response |
|----------|--------|----------|-------------|-------------------|----------|
| `/v1/hotels` | GET | REST (HTTP) | Retrieve hotels by IDs or destination | Query params: `hotelIDs[]`, `destinationId` | JSON array of hotels |
| `GetHotels` | RPC | gRPC | Retrieve hotels by IDs or destination | `GetHotelsRequest` | `GetHotelsResponse` |

**Request Body Parameters:**

**Sample REST API:**
```
GET /v1/hotels?hotelIDs=iJhz&destinationId=5432
```

**gRPC:**
```protobuf
message GetHotelsRequest {
  repeated string hotelIDs = 1;    // Array of hotel IDs to filter by
  uint64 destinationId = 2;        // Destination ID to filter by
}
```

**Sample Response Format:**
```json
{
	"hotels": [
		{
			"id": "iJhz",
			"destinationId": "5432",
			"name": "Beach Villas Singapore",
			"location": {
				"lat": 1.264751,
				"lng": 103.824006,
				"address": "8 Sentosa Gateway, Beach Villas, 098269",
				"city": "Singapore",
				"country": "SG"
			},
			"description": "Located at the western tip of Resorts World Sentosa, guests at the Beach Villas are guaranteed privacy while they enjoy spectacular views of glittering waters. Guests will find themselves in paradise with this series of exquisite tropical sanctuaries, making it the perfect setting for an idyllic retreat. Within each villa, guests will discover living areas and bedrooms that open out to mini gardens, private timber sundecks and verandahs elegantly framing either lush greenery or an expanse of sea. Guests are assured of a superior slumber with goose feather pillows and luxe mattresses paired with 400 thread count Egyptian cotton bed linen, tastefully paired with a full complement of luxurious in-room amenities and bathrooms boasting rain showers and free-standing tubs coupled with an exclusive array of ESPA amenities and toiletries. Guests also get to enjoy complimentary day access to the facilities at Asia’s flagship spa – the world-renowned ESPA.",
			"amenities": {
				"general": [
					"businesscenter",
					"indoor pool",
					"breakfast",
					"tub",
					"business center",
					"drycleaning",
					"outdoor pool",
					"childcare",
					"aircon",
					"wifi",
					"pool"
				],
				"room": [
					"iron",
					"tv",
					"coffee machine",
					"kettle",
					"hair dryer"
				]
			},
			"images": {
				"rooms": [
					{
						"link": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/2.jpg",
						"description": "Double room"
					},
					{
						"link": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/4.jpg",
						"description": "Bathroom"
					}
				],
				"site": [
					{
						"link": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/1.jpg",
						"description": "Front"
					}
				],
				"amenities": [
					{
						"link": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/0.jpg",
						"description": "RWS"
					},
					{
						"link": "https://d2ey9sqrvkqdfs.cloudfront.net/0qZF/6.jpg",
						"description": "Sentosa Gateway"
					}
				]
			},
			"bookingConditions": [
				"All children are welcome. One child under 12 years stays free of charge when using existing beds. One child under 2 years stays free of charge in a child's cot/crib. One child under 4 years stays free of charge when using existing beds. One older child or adult is charged SGD 82.39 per person per night in an extra bed. The maximum number of children's cots/cribs in a room is 1. There is no capacity for extra beds in the room.",
				"Pets are not allowed.",
				"WiFi is available in all areas and is free of charge.",
				"Free private parking is possible on site (reservation is not needed).",
				"Guests are required to show a photo identification and credit card upon check-in. Please note that all Special Requests are subject to availability and additional charges may apply. Payment before arrival via bank transfer is required. The property will contact you after you book to provide instructions. Please note that the full amount of the reservation is due before arrival. Resorts World Sentosa will send a confirmation with detailed payment information. After full payment is taken, the property's details, including the address and where to collect keys, will be emailed to you. Bag checks will be conducted prior to entry to Adventure Cove Waterpark. === Upon check-in, guests will be provided with complimentary Sentosa Pass (monorail) to enjoy unlimited transportation between Sentosa Island and Harbour Front (VivoCity). === Prepayment for non refundable bookings will be charged by RWS Call Centre. === All guests can enjoy complimentary parking during their stay, limited to one exit from the hotel per day. === Room reservation charges will be charged upon check-in. Credit card provided upon reservation is for guarantee purpose. === For reservations made with inclusive breakfast, please note that breakfast is applicable only for number of adults paid in the room rate. Any children or additional adults are charged separately for breakfast and are to paid directly to the hotel."
			]
		}
	]
}
```

## 6. How to Run the Test Cases

**Run All Tests:**
```bash
go test ./...
```

**Run Tests with Verbose Output:**
```bash
go test ./... -v
```

**Run Tests for Specific Package:**
```bash
go test ./internal/hotels/...
go test ./internal/suppliers/...
go test ./external/...
```

**Run Tests with Coverage:**
```bash
go test ./... -cover
```

**Run Tests for Specific Function:**
```bash
go test -run TestFunctionName ./package_path
```

## 7. Structure of this Application

**Structure at a glance**
```
hotelsDataMerge/
├── main.go                           # Application entry point
├── proto/                            # Protocol Buffer definitions
│   ├── hotelsdatamerge.proto         
│   └── google/api/                   
├── external/                         # External APIs (to get suppliers info)                  
├── internal/                         # Internal application logic
│   ├── hotels/                       # Hotel domain logic
│   └── suppliers/                    # Supplier domain logic
│       ├── fetcher/                  # Data fetching layer
│       ├── parser/                   # Data parsing layer
│       ├── merger/                   # Data merging layer
│       └── utils/                    # Utility functions
└── server/                           # gRPC and HTTP server
```

## 8. Main Design Patterns & Principles used

### 8.1. Factory Pattern
- **Location:** `internal/suppliers/parser/factory.go`
- **Purpose:** Creates appropriate parser instances based on supplier type
- **Implementation:** `DefaultParserFactory` with `CreateParser` method

### 8.2. Builder Pattern
- **Location:** `internal/suppliers/merger/hotel/builder.go`
- **Purpose:** Constructs merged hotel objects with fluent interface
- **Implementation:** `HotelBuilder` with method chaining for hotel construction

### 8.3. Dependency Injection
- **Location:** Throughout the application
- **Purpose:** Loose coupling between components
- **Implementation:** Logger and external dependencies injected via constructors

## 9. How is Dirty Data Being Cleaned

### 9.1. Data Normalization
- **String Trimming:** Removes leading/trailing whitespace from all string fields
- **Case Consistency:** Standardizes text formatting
- **Null Handling:** Converts null/empty values to appropriate defaults

### 9.2. Data Validation
- **Type Safety:** Ensures data types match expected schemas
- **Range Validation:** Validates coordinates, IDs, and numeric values
- **Required Fields:** Checks for mandatory data presence

### 9.3. Data Standardization
- **Country Codes:** Normalizes country representations
- **Address Formatting:** Standardizes address structure
- **Amenity Categorization:** Groups amenities into logical categories

## 10. Merging Techniques

### 10.1. Conflict Resolution Strategies

**`ID` Merging:**
- Prefers non-empty IDs over empty ones

**`Name` Merging:**
- Prefers longer/more descriptive name
- Prioritizes completeness over brevity

**`Location` Merging:**
- **`Coordinates`:** Prefers non-zero coordinates
- **`Address`:** Prefers longer, more detailed addresses
- **`City`:** Prefers non-empty city names
- **`Country`:** Prefers 2-letter country codes for consistency

**`Description` Merging:**
- Prefers the longer description for more comprehensive information

**`Amenities` Merging:**
- Combines amenities from all sources
- Removes duplicates

**`Images` Merging:**
- Aggregates images from all sources
- Maintains categorization (rooms, site, amenities)

**`Booking Conditions` Merging:**
- Combines all booking conditions
- Removes duplicate conditions

### 10.2. Merging Algorithm

**Step 1: Data Aggregation**
- Collects all hotel records from multiple suppliers
- Groups by hotel ID for matching

**Step 2: Conflict Detection**
- Identifies fields with different values
- Applies simple rules for resolution

**Step 3: Custom Selection**
- Uses quality metrics (length, completeness, format)
- Applies supplier-specific preferences

**Step 4: Result Construction**
- Builds final merged hotel object
- Ensures all fields are properly populated
- Validates final result

**Implementation in `internal/suppliers/merger/merge_hotels_data.go`:**
```go
func buildMergedHotel(existing, new hotels.Hotel) hotels.Hotel {
    hotelBuilder := mergerHotel.NewHotelBuilder(existing)
    
    hotelBuilder.WithID(existing.Id, new.Id)
    hotelBuilder.WithDestinationID(existing.DestinationId, new.DestinationId)
    hotelBuilder.WithName(existing.Name, new.Name)
    hotelBuilder.WithDescription(existing.Description, new.Description)
    hotelBuilder.WithLocation(existing.Location, new.Location)
    hotelBuilder.WithAmenities(existing.Amenities, new.Amenities)
    hotelBuilder.WithImages(existing.Images, new.Images)
    hotelBuilder.WithBookingConditions(existing.BookingConditions, new.BookingConditions)
    
    return hotelBuilder.Build()
}
```

# Areas for Improvement

As seen in the Sample Response Format, the `businesscenter` and `business center` are considered as 2 separate amenities. 

```go
      "amenities": {
				"general": [
					"businesscenter",
					"indoor pool",
					"breakfast",
					"tub",
					"business center",
					"drycleaning",
					"outdoor pool",
					"childcare",
					"aircon",
					"wifi",
					"pool"
				],
				"room": [
					"iron",
					"tv",
					"coffee machine",
					"kettle",
					"hair dryer"
				]
			}
```
In order to remove strings like `businesscenter` because `business center` is already present, there is a need to normalise the strings, and perform deduplication by defining a dictionary mapping. Another way is to calculate similarity of the strings using Levenshtein distance.

However, the logics for these 2 ways are not straightforward, thus this handling is not implemented.
