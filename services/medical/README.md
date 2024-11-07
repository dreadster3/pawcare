# Medical Records Service

## Description

The Medical Records Service is responsible for managing user profiles within the PawCare application. It handles operations such as creating, updating, retrieving, and deleting user profiles. This service ensures that user data is stored securely and can be accessed efficiently.

## Environment Variables

The following environment variables are used to configure the Medical Records Service:

- `DB_HOST`: The hostname of the database server.
- `DB_PORT`: The port number of the database server.
- `DB_NAME`: The URL of the database where user profiles are stored.
- `DB_USER`: The username used to connect to the database.
- `DB_PASSWORD`: The password used to connect to the database.
- `JWT_SECRET`: The secret key used for signing JSON Web Tokens.
- `LOG_LEVEL`: The level of logging detail (e.g., `info`, `debug`, `error`).

## Getting Started

To run the Medical Records Service locally using Docker Compose, follow these steps:

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/pawcare.git
   cd pawcare/services/profile
   ```

2. Set up environment variables:
   Create a `.env` file in the root of the project and add the necessary environment variables:

   ```env
   DB_HOST=db
   DB_USER=user
   DB_PASSWORD=password
   JWT_SECRET=jwt_secret
   ```

3. Start the service using Docker Compose:
   ```sh
   docker-compose up -d
   ```

This will build and start the Medical Records Service along with any other services defined in the `docker-compose.yaml` file.

## Contributing

If you would like to contribute to the Medical Records Service, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Make your changes and commit them with a descriptive message.
4. Push your changes to your fork.
5. Create a pull request to the main repository.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
