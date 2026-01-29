
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8081/swagger.json" -Method Get -ErrorAction Stop
    $json = $response | ConvertTo-Json -Depth 10
    if ($json -match "category_id") {
        Write-Host "Verification Success: category_id found in swagger.json"
    } else {
        Write-Error "Verification Failed: category_id not found"
    }
} catch {
    Write-Error $_.Exception.Message
}
