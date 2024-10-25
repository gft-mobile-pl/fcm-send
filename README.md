# fcm-send

Script for sending notification message using Firebase Cloud Messaging (FCM).

## Preparation
Generate `service-account-file.json` in Firebase.

## Usage
1. Copy `example_message.json` to `message.json`
2. In `message.json` change token to a token of your device. Change payload, if needed.
3. Check if there is file `service-account-file.json` in root folder
4. Run `./fcm-send -m message.json`
