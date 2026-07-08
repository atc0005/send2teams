// Copyright 2026 Adam Chalkley
//
// https://github.com/atc0005/send2teams
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

// Small CLI tool used to encode Microsoft Teams webhook URLs. `webhookenc` is
// intended for use by Nagios admins who may wish to encode a webhook URL for
// inclusion in a `Custom Object Variable` or a `User Macro` where
// sanitization would strip out required `&` characters used to separate URL
// query parameters in webhook URLs.
//
// See our [GitHub repo]:
//
//   - to review documentation (including examples)
//   - for the latest code
//   - to file an issue or submit improvements for review and potential
//     inclusion into the project
//
// [GitHub repo]: https://github.com/atc0005/send2teams
package main
