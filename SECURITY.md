# Security Policy

## Reporting a Vulnerability

If you discover a security vulnerability within `strinterp`, please follow these steps:

1.  **Do not disclose the vulnerability publicly** until it has been addressed.
2.  **Email the project maintainers** directly at [ujjwal.kumar1@ibm.com] with a detailed description of the vulnerability, including:
    -   Steps to reproduce the vulnerability.
    -   Affected versions (if known).
    -   Any potential impact.
3.  **Allow reasonable time** for the maintainers to investigate and develop a fix.
4.  **Cooperate** with the maintainers in testing the fix and coordinating the release.

We take security seriously and appreciate your efforts to responsibly disclose any issues. We will acknowledge your contribution in the release notes unless you request otherwise.

## Supported Versions

Currently, we provide security updates for the latest released version of `strinterp`. We encourage users to always use the most recent version.

## Security Considerations

While `strinterp` itself is designed to be secure, be cautious when using it with user-supplied input, particularly when enabling features like custom formatters or function calls (if implemented in the future).  Always sanitize and validate user input to prevent potential code injection or other security risks.

## Out of Scope

The following are considered out of scope for security reports:

-   Vulnerabilities in dependencies (report these to the respective projects).
-   Issues related to insecure usage or misconfiguration of the library.
-   Attacks requiring physical access to the system.
-   Social engineering attacks.

Thank you for helping to keep `strinterp` secure!
