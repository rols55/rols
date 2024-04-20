"use client";
import { useState } from "react";
import { useRouter } from "next/navigation";
import Cookies from 'js-cookie';

function LoginPage() {
    const [formData, setFormData] = useState({
        username: "",
        password: "",
    });

    const handleChange = (event) => {
        const { name, value } = event.target;
        setFormData((prevFormData) => ({
            ...prevFormData,
            [name]: value,
        }));
    };


    const [formErrors, setFormErrors] = useState({
      invalidCredentials: ""
    });

    /*
    const getErr = (name, value) => {
      switch (name) {
          case "invalidCredentials":
              if (!value.trim()) {
                  return "This field is required";
              }
              break;
      }
      return "";
  };*/


  const router = useRouter();

  const handleLogin = async (e) => {
    e.preventDefault(); // Prevent default form submission

    const payload = new FormData();

    for (const key in formData) {
        payload.append(key, formData[key]);
    }
    const options = {
        method: "POST",
        body: payload,
    };

    try {
        const response = await fetch(
            `http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/login`,
            options
        );
        if (response.ok) {
            const data = await response.json();
            Cookies.set(data.CookieName, data.SessionToken , { expires: data.ExpirationTime, sameSite: 'None', secure: true}); // Cookie expires in 7 days
            router.push("/");
            router.refresh();
        } else {
            const statusMsg = await response.text();
            console.log(statusMsg);
            if (statusMsg.includes("Wrong password") || statusMsg.includes("doesn't exist")) {
              setFormErrors((prevErrs) => ({
                  ...prevErrs,
                  invalidCredentials: "Invalid username or password. Please try again.",
              }));
          }
        }
    } catch (error) {
        console.error(error);
    }
  };

  return (
    <div className="login">
      <form onSubmit={handleLogin}>
        <label>
          Username:
          <input
            type="text"
            name ="username"
            value={formData.username} onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
          Password:
          <input
            type="password"
            name ="password"
            value={formData.password} onChange={handleChange}
            required
          />
        </label>

        <div className="feedback-error">
                {formErrors.invalidCredentials}
            </div>

        <br />
        <button type="submit">Log In</button>
      </form>
    </div>
  );
}

export default LoginPage;