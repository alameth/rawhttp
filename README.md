NAME
====

rawhttp -- HTTP Testing Client

SYNOPSIS
========

`rawhttp` [_options_] _servername_[:_port_]

DESCRIPTION
===========
**rawhttp** is a test client for HTTP and HTTPS. It's a YAFIYGI client (you asked
for it, you got it) that does what it's told and stays out of the way. **rawhttp**
makes no effort to ensure correctness or handle responses unless told to do so.
**rawhttp** will never issue more than one request or listen for more than one
response; interpreting the response is the sole responsibility of the caller.

With no options, **rawhttp** acts as follows:
- Read a complete HTTP Request (header and body) from stdin;
- Connect to the specified server on port 80;
- Write the Request to the server;
- Read the response from the server;
- Write the response header to stdout.

The Request header is normalized to CRLF line termination before writing,
but the body is unchanged in order to allow binary post data.

Options are available to use TLS (HTTPS), read the request from one or more
files, set the port number, and to automate some simple elements in the Request
header. If the _server_ is recognized as a URL, then the scheme, host name, and
port number are extracted.

OPTIONS
=======
<DL>
<DT>-b filename</DT>
<DD>Read the request body from the specified file.
If this option is present, the <TT>-h</TT> option is required as well.
If you really want to send a reqiest that has a body and no header,
use the <TT>-m</TT> option.</DD>
<DT>-e</DT>
<DD>Set the exit status code based on the HTTP response.
If the response is 2xx, then the exit status code is set to 0.
If a network error occurs, the exit status code is set to 1.
If the response code is 3xx, 4xx, or 5xx, then the exit status
code is set to 3, 4, or 5, respectively.</DD>
<DT>-h filename</DT>
<DD>Read the request header from the specified file.
If this option is present and the <TT>-b</TT> option is not, no body is sent.
This option cannot be combined with the `-m` option.</DD>
<DT>-i filename</DT>
<DD>Read the input message (request) from the specified file, rather than from stdin.
This option cannot be combined with the `-h` or `-b` options.</DD>
<DT>-o filename</DT>
<DD>Write the output (response) to the specified file.
The file will include the complete response header and body, without parsing
or interpretation.</DD>
<DT>-v</DT>
<DD>Verbose output.</DD>
<DT>-S servername</DT>
<DD>Manually the Server Name in SNI. By default it is set to the server name.
If the name is set to - (the minus sign) then generation of the SNI will be
supressed entirely.</DD>
<DT>-T</DT>
<DD>Use TLS. The port number default is changed to 443 and SNI is set to the
server name. Use the -S option to change the SNI name.</DD>
</DL>

EXIT STATUS
===========
**rawhttp** may return 