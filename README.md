Just a little tool to generate a unix-style SHA-512 password hash, with a random salt.

To build it, you'll need to have https://golang.org/ installed.

    $ git clone https://github.com/yobert/joelcrypt
    $ cd joelcrypt
    $ go build

Use it like:

    $ echo -n "secret" | ./joelcrypt 
    $6$5E6gEtI1in70EjNT$X498k7VrPKms.ncNsDFRYv/UkT580fyPu05xCtrklW0W2ExoPU93gTbCRIcvQlE7lE.Jp8mQhIEf7qc6XB.PP0

Make sure to remember the -n or it will include the newline in your password.
