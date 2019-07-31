/*
Package db is a wrapper for the sql package.

The main aim of the wrapper is to make sure that with normal actions, the
context should be consider in order to better control the flow of the code.

If there is no need for the context, then passing nil is enough.
Please make sure that it is the right call before hand.
*/
package db
