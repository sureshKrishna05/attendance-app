class ApiConfig {
  // Use http://10.0.2.2:8080 for Android Emulator pointing to localhost.
  // Use http://localhost:8080 for iOS Simulator or Web.
  // Change to your server IP for physical devices.
  static const String baseUrl = 'http://127.0.0.1:8080/api/v1'; // Assuming web/desktop for now, user can change this.
  static const String publishableKey = 'pk_live_12345'; // Matching the docker-compose .env
}
