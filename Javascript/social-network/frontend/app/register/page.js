"use client";
import { useState } from "react"
import { useRouter} from "next/navigation"

export default function Contact() {

    const router = useRouter();

    const [formData, setFormData] = useState({
        username: "",
        gender: "",
        age: "",
        firstName: "",
        lastName: "",
        nickName: "",
        email: "",
        aboutMe: "",
        password: "",
        passwordConfirm: "",
    });

    const [formErrors, setFormErrors] = useState({
        username: "",
        email: "",
        usernameExist: ""
    });

    const [imageFile, setImageFile] = useState(null);

    const handleImageChange = (event) => {
        const file = event.target.files[0];
        setImageFile(file);
      };

    const handleChange = (event) => {
        const { name, value } = event.target;
        setFormData((prevFormData) => ({
            ...prevFormData,
            [name]: value,
        }));

        const err = getErr(name, value);
        setFormErrors((prevErrs) => ({
            ...prevErrs,
            [name]: err,
        }));
    };

    const getErr = (name, value) => {
        switch (name) {
            case "username":
                if (!value.trim()) {
                    return "This field is required";
                }
                break;
            case "email":
                if (
                    !/^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/.test(value)
                ) {
                    return "Not a valid email address";
                }
                break;
            case "password":
                if (value.length < 6) {
                    return "Password is too short";
                }
                break;
            case "passwordConfirm":
                if (formData.password !== value) {
                    return "Passwords don't match";
                }
                break;
            case "usernameExist":
                if (formData.password !== value) {
                    return "Passwords don't match";
                }
                break;
        }
        return "";
    };

    const handleSubmit = async (event) => {
        event.preventDefault();

        const payload = new FormData();

        for (const key in formData) {
            payload.append(key, formData[key]);
        }

        if (imageFile) {
            payload.append("image", imageFile);
        }

        const options = {
            method: "POST",
            body: payload,
        };

        try {
            const response = await fetch(
                `http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/register`,
                options
            );
            if (response.ok) {
                router.push("/login");
            } else {
                const statusMsg = await response.text();
                console.log(statusMsg);
                if (statusMsg.includes("already taken")) {                    
                    setFormErrors((prevErrs) => ({
                        ...prevErrs,
                        usernameExist: "This username is already in use",
                    }));
                }
            }
        } catch (error) {
            console.error(error);
        }
    };

  return (
    <div className="register">
      <h1>Register</h1>
        <form id="signupForm" onSubmit={handleSubmit}>
          <div>
            <label>Choose username</label>
            <input type="text" className={`form-control ${formErrors.username ? "is-invalid" : ""
                }`} id="signupUsername" name="username"  value={formData.username} onChange={handleChange} />
          </div>
          <div className="feedback-error">
                                    {formErrors.username}
                                </div>

          <div>
            <label>Sex</label>
            <select id="signupSex" name="gender"  value={formData.gender} onChange={handleChange} required>
                        <option value="" disabled>Select gender</option>
                        <option value="male">Male</option>
						<option value="female">Female</option>
						<option value="other">Other</option>
				</select>          </div>

          <div>
			<label>Birthday</label>
            <input type="date" id="birthday" name="age"  value={formData.age} onChange={handleChange} required/>
          </div>

          <div>
          <label>First Name</label>
          <input type="text" id="signupfirstname" name="firstName"  value={formData.firstName} onChange={handleChange} required/>
          </div>

          <div>
          <label>Last Name</label>
          <input type="text" id="signupLastname" name="lastName"  value={formData.lastName} onChange={handleChange} required/>
          </div>

          <div>
          <label>Nickname (optional)</label>
          <input type="text" id="nickName"   name="nickName" value={formData.nickName} onChange={handleChange}  />
          </div>

          <div>
          <label>E-mail</label>
          <input type="email" id="signupEmail" name="email" className={`form-control ${
                                        formErrors.email ? "is-invalid" : ""
                                    }`}  value={formData.email} onChange={handleChange} required />
          </div>

          <div className="feedback-error">
                                    {formErrors.email}
                                </div>

          <div>
          <label>Add picture (optional)</label>
          <input type="file"  id="imageInput" name="image"  onChange={handleImageChange} />
          </div>



          <div>
          <label>About me (optional)</label>
          <input type="text"   name="aboutMe" value={formData.aboutMe} onChange={handleChange} />
          </div>

          <div>
          <label>Password</label>
          <input type="password" className={`form-control ${
                                                formErrors.password
                                                    ? "is-invalid"
                                                    : ""
                                            }`} id="signupPassword" name="password"  value={formData.password} onChange={handleChange} required/>
          </div>

          <div className="feedback-error">
                                            {formErrors.password}
                                        </div>

          <div>
          <label>Re-enter password</label>
          <input type="password"                                             className={`form-control ${
                                                formErrors.passwordConfirm
                                                    ? "is-invalid"
                                                    : ""
                                            }`} id="passwordConfirm" name="passwordConfirm"  value={formData.passwordConfirm} onChange={handleChange} required/>
          </div>

          <div className="feedback-error">
                                            {formErrors.passwordConfirm}
                                        </div>

            <div className="feedback-error">
                {formErrors.usernameExist}
            </div>


          <input type="submit" id="submitSignupForm" value="Sign up"/>
        </form>
    </div>
  )
}

