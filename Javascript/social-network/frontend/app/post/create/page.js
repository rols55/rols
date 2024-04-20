"use client";
import React, { useEffect, useState } from "react"
import { useRouter} from "next/navigation"



function Create() {

    const router = useRouter();

    const [formData, setFormData] = useState({
        title: "",
        text: "",
        privacy: "",
	    followers: ""
    });

    const [formErrors, setFormErrors] = useState({
        title: "",
        text: "",
        privacy: "",
        followers: ""
    });

    useEffect(() => {
        const fetchFollowers = async () => {
          try {
            const response = await fetch(`http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/followers`);
            if (response.ok) {
              const data = await response.json();
              console.log(data.followers);
              const followers = data.followers
                .filter((follower) => follower.accepted === true)
                .map((follower) => follower.username);
              setFollowersArray(followers); // Update followersArray state
            } else {
              const statusMsg = await response.text();
              console.log(statusMsg);
              // Handle error
            }
          } catch (error) {
            console.error(error);
            // Handle error
          }
        };
    
        fetchFollowers();
      }, []); // Run only once when the component mounts

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
/*
    const getErr = (name, value) => {
        switch (name) {
            case "title":
                if (!value.trim()) {
                    return "This field is required";
                }
                break;
            case "text":
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
        }
        return "";
    };
*/
    const handleCheckboxChange = (event) => {
        const { name, checked } = event.target;
      
        
      
        setFormData((prevFormData) => {
          if (checked) {
            // Add follower to the formData
            return {
              ...prevFormData,
              followers: [...prevFormData.followers, name],
            };
          } else {
            // Remove follower from the formData
            return {
              ...prevFormData,
              followers: prevFormData.followers.filter(follower => follower !== name),
            };
          }
        });
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
                `http://${process.env.NEXT_PUBLIC_API_HOST_CLIENT_SIDE}:${process.env.NEXT_PUBLIC_API_PORT_CLIENT_SIDE}/api/post/create`,
                options
            );
            if (response.ok) {
                router.push("/");
                const data = await response.json();
            } else {
                const statusMsg = await response.text();
                console.log(statusMsg);
            }
        } catch (error) {
            console.error(error);
        }
    };

  return (
    <div>
      <h1>Create new post</h1>
        <form id="signupForm" onSubmit={handleSubmit}>
          <div>
          <label>Title</label>
          <input type="text" id="Title" name="title"  value={formData.title} onChange={handleChange} required/>
          </div>
          <div>
          <label>Text</label>
          <input type="text" id="Text" name="text"  value={formData.text} onChange={handleChange} required/>
          </div>

          <div>
          <label>Add picture (optional)</label>
          <input type="file"  id="imageInput" name="image"  onChange={handleImageChange} />
          </div>

          <div>
          <label>Privacy</label>
            <select id="Privacy" name="privacy" value={formData.privacy} onChange={handleChange} required>
              <option value="">Select Privacy</option>
               <option value="public">Public</option>
              <option value="private">Private</option>
            </select>
          </div>
          <div className="feedback-error">
                                    {formErrors.privacy}
            </div>

            {formData.privacy === "private" && (
  <div>
    <label>Followers allowed</label>
    {followersArray.length > 0 && (
  <>
    {followersArray.map(follower => (
      <div key={follower}>
        <input 
          type="checkbox" 
          id="followers" 
          name={follower} 
          value={follower} 
          checked={formData.followers.includes(follower)} 
          onChange={handleCheckboxChange} 
        />
        <label htmlFor={follower}>{follower}</label>
      </div>
    ))}
  </>
)}
</div>
)}

          <input type="submit" id="submitSignupForm" value="Sign up"/>
            </form>
            </div>
  )

}

export default Create;