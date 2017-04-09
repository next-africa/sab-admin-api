/**
 * Created by pdiouf on 2017-04-08.
 */
fetch("httcountries/"+formPayload.countryCode+"/universities/", {method: "POST", body: JSON.stringify({formPayload})})
    .then((response) => response.json())
    .then((responseData) => {
        console.log(
            "POST Response",
            "Response Body -> " + JSON.stringify(responseData.body)
        )
    });