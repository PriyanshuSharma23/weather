# Weather CLI Tool

<!--toc:start-->

- [Weather CLI Tool](#weather-cli-tool)
  - [Features](#features)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Output](#output)
  - [Example Output](#example-output)
  - [Error Handling](#error-handling)
  - [Customization](#customization)
  - [Dependencies](#dependencies)
  - [License](#license)
  - [Author](#author)
  <!--toc:end-->

This is a simple CLI tool that fetches and displays the current weather information for a specified location using the WeatherAPI. It provides details such as temperature, feels like temperature, humidity, wind speed, and precipitation chance. If no location is provided, the tool defaults to a predefined location set via an environment variable.

## Features

- Fetches real-time weather data from WeatherAPI.
- Displays the temperature, humidity, wind speed, and precipitation chances.
- Uses color codes and emojis to represent different temperature ranges.
- Falls back to a default location set via an environment variable if no location is specified.

## Prerequisites

- **Go**: This tool is written in Go, so you will need Go installed to build and run it.
- **WeatherAPI Key**: You need an API key from [WeatherAPI](https://www.weatherapi.com/) to use this tool. Replace the placeholder `API_KEY` in the code with your actual API key.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/weather-cli.git
   cd weather-cli
   ```

2. Set up your environment variables:

   Set the `DEFAULT_WEATHER_LOCATION` environment variable for the default location. Example:

   ```bash
   export DEFAULT_WEATHER_LOCATION="New York"
   ```

   Replace `"New York"` with the city you want to set as your default location.

3. Replace the placeholder API key:

   Open the `main.go` file and replace the placeholder `API_KEY` value with your actual WeatherAPI key:

   ```go
   const API_KEY  = "your_actual_api_key"
   ```

4. Build the application:

   ```bash
   go build -o weather
   ```

5. Run the application:

   ```bash
   ./weather [location]
   ```

   Example usage:

   ```bash
   ./weather "San Francisco"
   ```

   If no location is provided, the tool will use the default location set in the environment variable `DEFAULT_WEATHER_LOCATION`.

## Usage

To fetch weather information, run the CLI tool with the desired location as an argument:

```bash
./weather "Los Angeles"
```

If you don't provide a location, it will fall back to the default location set via the `DEFAULT_WEATHER_LOCATION` environment variable.

For help, you can display usage instructions:

```bash
./weather --help
```

## Output

The tool displays the following weather information:

- **Temperature**: The current temperature at the location.
- **Feels Like**: The perceived temperature based on wind and humidity.
- **Humidity**: The percentage of humidity in the air.
- **Wind Speed**: The speed of the wind.
- **Precipitation Chance**: The chance of rain or precipitation.

The temperature is also displayed with colors and emojis based on the current value:

- **< 20°C**: 🥶 (Blue color)
- **20°C - 30°C**: 😎 (Green color)
- **30°C - 40°C**: 🥺 (Yellow color)
- **> 40°C**: 🥵 (Red color)

## Example Output

```bash
🥶 15.00°C, feels like 13.00°C
📍 San Francisco, California, United States
```

## Error Handling

If the API request fails or the location cannot be found, an error message will be displayed in red.

## Customization

- **Default Location**: You can set your default location by setting the `DEFAULT_WEATHER_LOCATION` environment variable.
- **Temperature Ranges**: You can customize the temperature ranges and associated emojis in the `weather.Print()` method.
- **Error Handling**: Error handling is managed with the `printError()` function.

## Dependencies

This project uses the following dependencies:

- `logx`: A small logging utility for colorful output in the terminal.
- `net/http`: For making API requests.
- `encoding/json`: For decoding the JSON response from WeatherAPI.

Install dependencies using Go modules:

```bash
go mod tidy
```

## License

This project is licensed under the MIT License. Feel free to use and modify it as needed.

## Author

- **Priyanshu Sharma**
