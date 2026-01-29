
try {
    Write-Host "Creating Product..."
    $postParams = @{ name="Soto"; price=3000; stock=100; category_id=1 } | ConvertTo-Json
    $response = Invoke-RestMethod -Uri "http://localhost:8081/api/product" -Method Post -Body $postParams -ContentType "application/json" -ErrorAction Stop
    Write-Host "Create Response:"
    Write-Host ($response | ConvertTo-Json)

    Write-Host "`nUpdating Product 1..."
    $putParams = @{ name="Soto update"; price=4000; stock=500; category_id=1 } | ConvertTo-Json
    $responsePut = Invoke-RestMethod -Uri "http://localhost:8081/api/product/1" -Method Put -Body $putParams -ContentType "application/json" -ErrorAction Stop
    Write-Host "Update Response:"
    Write-Host ($responsePut | ConvertTo-Json)
    
    Write-Host "`nVerification Success"
} catch {
    Write-Error $_.Exception.Message
}
