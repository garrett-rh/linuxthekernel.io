import { useState, useEffect } from "react";

const CarTax = () => {
    const [value, setValue] = useState(0);
    const [locality, setLocality] = useState("");
    const [tax, setTax] = useState(0)
    const [localities, setLocalities] = useState([])

    useEffect( () => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const requestOptions = {
            method: "GET",
            headers: headers,
        }

        fetch(`/api/car_tax/localities`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                setLocalities(data);
            })
            .catch(err => {
                console.error(err);
            })

    }, []);


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
                setTax(result.taxes);
            })
            .catch((error) => {
                console.error(error);
            });
    };

    return (
        <div>
            <p className="fs-1 text-center">Vehicle Personal Property Tax Calculator</p>
            <hr />
            <form onSubmit={SubmitCarInfo}>
                <div className="mb-3">
                    <label htmlFor="CarValue" className="form-label">Car Value in Dollars</label>
                    <input
                        type="number"
                        className="form-control"
                        id="CarValue"
                        value={value}
                        onChange={(e) => setValue(e.target.valueAsNumber)}
                        required="true"
                    />
                </div>
                <div className="mb-3">
                    <select
                        className="form-select"
                        value={locality}
                        onChange={(e) => setLocality(e.target.value)}
                        required="true"
                    >
                        <option value="" disabled="true">Select a Locality</option>
                        {localities.map((m) => (
                            <option key={m.name} value={m.name}>{m.name}</option>
                        ))}
                    </select>
                </div>
                <button type="submit" className="btn btn-primary">Submit</button>
            </form>
            <div className="card mt-4 text-center">
                <div className="card-body">
                    <h5 className="card-title">Your Total Tax Burden</h5>
                    <p className="card-text fs-4 fw-bold text-primary">
                        ${tax.toFixed(2)}
                    </p>
                    <p className="card-text">This is the total tax over 12 months.</p>
                </div>
            </div>
        </div>
    );
};

export default CarTax;
