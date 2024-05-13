# sendEmailViaGolang
This is a repository that is meant to be a teaching tool on how to send email from your google account using golang.

# Setup

1. Login to your google account in a webbrowser.

2. Go to "Manage your google account".

3. Click "Security" on the left.

4. Scroll down to "2-Step Verification" and you will need to make sure it is turned on.

5. Now that you are in 2-Step verification page you may need to set your account up with a phone number.

6. Now that 2-Step verification is on you will need to go back to "Manage your google account".

7. Search up "app pass" and you should see somthing called "App passwords", click on it.

8. Near the bottom of the page you will see somthing that says "App Name" and a text field below that.

9. You can put whatever you want in the text field, just put somthing that you will remember.

10. You will now get a popup that says "Generated app password" and it will give you text that looks like this 'xxxx xxxx xxxx xxxx'.

11. You can copy that code and go to your project directory and once you are in the "src" directory run this command.

```
mv example.env .env
```

12. Now open the ".env" file with the text editor of your choice and paste the code that you copied into the "EMAIL_PASSWORD" section.

```
nvim .env
```

13. You can now continue to the next section.

# Sending an email

1. Make sure you are in the "src" directory and have completed the "Setup" section.

2. Open the ".env" file with the text editor of your choice.

```
nvim .env
```

3. Fill out the "EMAIL_USERNAME" section with the email that you used in the setup section Ex. "john.doe@gmail.com".

4. Fill out the "TO_EMAIL" section with the email address that you are trying to send an email to Ex. "emailToSendTo@gmail.com".

5. You can now run "go run main.go <subject> <body>" (Make sure to use double quotes if you are going to put spaces in the subject or body).

6. Alternativly you can edit the main.go file and find the line with "!!!" on both sides (Line 20). The you can uncomment the 2 lines under the line. Then fill out the subject and body in the file (you can now run the file with just "go run main.go").
