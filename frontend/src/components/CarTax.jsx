import { useState } from "react";

const CarTax = () => {
    const [value, setValue] = useState();
    const [locality, setLocality] = useState("");

    const SubmitCarInfo = (event) => {
        event.preventDefault(); // Prevent form submission from reloading the page

        const data = {
            value: value,
            locality: locality,
        };

        const requestOptions = {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        };

        fetch("/api/car_tax/calculate", requestOptions)
            .then((response) => response.json())
            .then((result) => {
                console.log("Response from API:", result);
            })
            .catch((error) => {
                console.error("Error sending data:", error);
            });
    };

    return (
        <form onSubmit={SubmitCarInfo}>
            <div className="mb-3">
                <label htmlFor="CarValue" className="form-label">Car Value in Dollars</label>
                <input
                    type="number"
                    className="form-control"
                    id="CarValue"
                    value={value}
                    onChange={(e) => setValue(e.target.valueAsNumber)}
                />
            </div>
            <div className="mb-3">
                <select
                    className="form-select"
                    value={locality}
                    onChange={(e) => setLocality(e.target.value)}
                >
                    <option value="" disabled>Select a Locality</option>
                    <option value="Arlington">Arlington County</option>
                </select>
            </div>
            <button type="submit" className="btn btn-primary">Submit</button>
        </form>
    );
};

export default CarTax;
