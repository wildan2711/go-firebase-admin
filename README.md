go-firebase-admin
==============================

[![Go Report Card](https://goreportcard.com/badge/github.com/acoshift/go-firebase-admin)](https://goreportcard.com/report/github.com/acoshift/go-firebase-admin)
[![GoDoc](https://godoc.org/github.com/acoshift/go-firebase-admin?status.svg)](https://godoc.org/github.com/acoshift/go-firebase-admin)
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)

Table of Contents
==============================

 * [Overview](#overview)
 * [Installation](#installation)
 * [Features](#features)
 * [To-Do List](#to-do-list)
 * [Documentation](#documentation)
 * [Examples](#examples)
 * [License](#license)

Overview
==============================
Firebase Admin SDK for Golang

On Wednesday, May 17, 2017 [Google announced at Google IO][1] : Open sourcing the Firebase SDKs.
But for now, there is no official Admin SDK for Golang, only Java, Node and Python Admin SDKs.

So welcome go-firebase-admin SDK :)

> Note
```
If you decide to use this unofficial SDK still in development, 
please use any package manager to fix version, there will be a lot of breaking changes.
```

Installation
------------

Install the package with go:

    go get github.com/acoshift/go-firebase-admin

Features
--------

This go-firebase-admin SDK supports the following functions :

- Authentication
  * CreateCustomToken : [Generate JSON Web Tokens (JWTs) on your server][3], pass them back to a client device, and then use them to authenticate via the signInWithCustomToken() method.
  * VerifyIDToken : [verify the integrity and authenticity of the ID token][4] and retrieve the uid from it.
- User Management API
  * GetUser : fetching the profile information of users by their uid
  * GetUsers : fetching list of profile information of users by their uid
  * GetUserByEmail : fetching the profile information of users by their email
  * GetUsersByEmail : fetching list of profile information of users by their email
  * ListUsers : fetching the profile information of users
  * CreateUser : create a new Firebase Authentication user
  * UpdateUser : modifying an existing Firebase user's data.
  * DeleteUser : deleting existing Firebase user by uid
  * SendPasswordResetEmail : send password reset for the given user
  * VerifyPassword : verifies given email and password

- Realtime Database API
  * not documented
  
- Cloud Messaging API
  * SendToDevice : Send Message to individual device
  * SendToDevices : Send multicast Message to a list of devices
  * SendToDeviceGroup : Send Message to a device group
  * SendToTopic : Send Message to a topic
  * SendToCondition : Send a message to devices subscribed to the combination of topics
  * SubscribeDeviceToTopic : subscribe a device to a topic
  * SubscribeDevicesToTopic : subscribe devices to a topic
  * UnSubscribeDeviceFromTopic : Unsubscribe a device to a topic
  * UnSubscribeDevicesFromTopic : Unsubscribe devices to a topic

To-Do List
----------

- [ ] update documentation
- [ ] add examples
- [ ] add tests

Documentation
-------------

You can find more details about go-firebase-admin on [godoc.org][2].


* [Firebase Setup Guide](https://firebase.google.com/docs/admin/setup/)
* [Firebase Database Guide](https://firebase.google.com/docs/database/admin/start/)
* [Firebase Authentication Guide](https://firebase.google.com/docs/auth/admin/)
* [Firebase Cloud Messaging Guide](https://firebase.google.com/docs/cloud-messaging/admin/)
* [Firebase Cloud Messaging Server](https://firebase.google.com/docs/cloud-messaging/server)
* [Firebase Release Notes](https://firebase.google.com/support/releases)


Examples
--------

You need a service_account.json file, if you don't have an admin SDK service_account.json, please [check this guide]
(https://firebase.google.com/docs/admin/setup#add_firebase_to_your_app)  

Initialize Firebase Admin SDK

```go
package main

import (
  "io/ioutil"

  "github.com/acoshift/go-firebase-admin"
)

func main() {
  // Init App with service_account
  serviceAccount, _ := ioutil.ReadFile("service_account.json")
  firApp, err := admin.InitializeApp(context.Background(), admin.AppOptions{
    ServiceAccount: serviceAccount,
    ProjectID:      "YOUR_PROJECT_ID",
    DatabaseURL:    "YOUR_DATABASE_URL",
    APIKey:         "YOUR_API_KEY",
  })

  if err != nil {
    panic(err)
  }

  // Firebase AUth
  firAuth := firApp.Auth()

  // Firebase Database
  firDatabase := firApp.Database()

  // FCM
  firFCM := firApp.FCM()

  // SendToDevice
  resp, err := firFCM.SendToDevice(context.Background(), "mydevicetoken",
		admin.Message{Notification: admin.Notification{
			Title: "Hello go firebase admin",
			Body:  "My little Big Notification",
			Color: "#ffcc33"},
		})
	if err != nil {
		panic(err)
	}

  // SendToDevices
  resp, err := firFCM.SendToDevices(context.Background(), []string{"mydevicetoken"},
		admin.Message{Notification: admin.Notification{
			Title: "Hello go firebase admin",
			Body:  "My little Big Notification",
			Color: "#ffcc33"},
		})
	if err != nil {
		panic(err)
	}

  // SubscribeDeviceToTopic
  resp, err := firFCM.SubscribeDeviceToTopic(context.Background(), "mydevicetoken", "/topics/gofirebaseadmin")
  // it's possible to ommit the "/topics/" prefix
  resp, err := firFCM.SubscribeDeviceToTopic(context.Background(), "mydevicetoken", "gofirebaseadmin")
	if err != nil {
		panic(err)
	}

  // UnSubscribeDeviceFromTopic
  resp, err := firFCM.UnSubscribeDeviceFromTopic(context.Background(), "mydevicetoken", "/topics/gofirebaseadmin")
  // it's possible to ommit the "/topics/" prefix
  resp, err := firFCM.UnSubscribeDeviceFromTopic(context.Background(), "mydevicetoken", "gofirebaseadmin")
	if err2 != nil {
		panic(err)
	}
  
}
```

Licence
-------

MIT License

Copyright (c) 2016 Thanatat Tamtan

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

[1]: https://opensource.googleblog.com/2017/05/open-sourcing-firebase-sdks.html
[2]: https://godoc.org/github.com/acoshift/go-firebase-admin
[3]: https://firebase.google.com/docs/auth/admin/create-custom-tokens
[4]: https://firebase.google.com/docs/auth/admin/verify-id-tokens
[5]: https://firebase.google.com/docs/auth/admin/manage-users

