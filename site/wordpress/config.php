<?php
$servername = "localhost";
$database = "desafio02";
$username = "redteamsw";
$password = "y0um4yh4ckm3";
// Create connection
$conn = mysqli_connect($servername, $username, $password, $database);
// Check connection
if (!$conn) {
    die("Connection failed: " . mysqli_connect_error());
}
echo "Connected successfully";
mysqli_close($conn);
?>
