INSERT INTO users (id,uuid,username,firstname, lastname, sex, birthday, email, public, nickname, aboutme, image,password) VALUES
  (4,"900e0cd1-2529-4d2d-8bfc-99b061acd8e2","Ninja777", "John", "Smith", "male", "1970-01-01","ninja777@mail.com",true ,"","","defaultImage.png","$2a$10$FYqIsLxsRXofqFAUT3hwlOePon.xsa5VQV3S8WcS3BSK2qikn8aIe"),
  (5,"12ef5526-049a-4857-b46d-fc90db00cb92","DigitalNomad", "Jane", "Johnson", "female","1989-11-03","digitalnomad@email.com",false,"","","defaultImage.png","$2a$10$sr1XVHaiu0ijvJoG67Jkj.BrKmjo1pcY6SlS5z7ah.LrEFvXnUHFG"),
  (6,"945c7df5-88a2-4c58-bd6d-d92e04c6e282","CyberPirate", "Michael", "Williams", "male", "1995-03-15", "cyberpirate@email.com",true,"","","defaultImage.png","$2a$10$ic2sOnj0cIF88fjCPSi4m.exV3EvgKqS9ZxJ48E2pHKtltuOUeDJm"),
  (7,"837fcac6-c188-49f3-8c1c-8f99dcfef5b3","ForumJunkie", "David", "Brown", "male", "1999-07-16","forumjunkie@email.com",false,"","","defaultImage.png","$2a$10$.XkbCa3Pq7RSRlZv/ADMcenDgK3GUZ45X9HB2rmPi4q0ITNSVHVbq"),
  (8,"de98c4ff-387c-474b-af2e-8dad191915fb","tester", "Emily", "Jones", "female", "1981-08-10", "test@mail.com",true,"","","defaultImage.png","$2a$10$KwnFxK3MbvDRxNgJq/xroeN9JaZ1PuJvppcbaMFNo15JoAy7VD.Oq"),
  (9,"d644ae49-0a31-4409-afea-cd8c5c3006de","test" , "Olivia", "Wilson" , "female", "1991-02-10","test@test.com",false,"","","defaultImage.png","$2a$10$EPlZBo8BlqSHEqWeRaL/sefM17DxGRgJd2Af8Jyttbdsqmge7m2Yq"),
  (10,"1c8e81c8-95d7-4d5a-a248-9d042f8e974c","DarthVader", "Snorty", "Jamboree", "male", "1973-09-21", "vader@email.com",true,"","","defaultImage.png","$2a$10$EPlZBo8BlqSHEqWeRaL/sefM17DxGRgJd2Af8Jyttbdsqmge7m2Yq")
;

INSERT INTO posts (id, user_id, author, title, text, image, privacy, followers, creation_date, group_id) VALUES
  (3, "900e0cd1-2529-4d2d-8bfc-99b061acd8e2", "Ninja777", "Does your cats like listening to music?", "It's good to know what your cats love and do it for them. It's going to make them lovely being around with all the time because they will make it easy for have fun with them.   I found out my cats loves listening to musics like rock band songs. The way they pay attention to the songs when it plays helped me to figure it out. One of them wags her talk to the music beats.   Do you know if your cats loves music?", "", "public", "", '2023-06-08 10:43:39.8098142+03:00', 0),
  (4, "945c7df5-88a2-4c58-bd6d-d92e04c6e282", "CyberPirate", "Frog attempting escape", "Last night my 6 month old american bullfrog tried multiple time to jump through the mesh lid his tank. The only way i could stop him was by turning on his light so that he would hide. Any ideas to stop this behavior without leaving the light on 24/7?", "", "public", "", '2023-06-08 11:43:39.8098142+03:00', 0),
  (6, "d644ae49-0a31-4409-afea-cd8c5c3006de", "test", "I have a parrot but would like a lap dog too. What are some good breeds?", "Ive been spending the last couple weeks researching dogs with lower prey drives, how they deal with smaller animals, stuff like that as well as some needs that would suit my life like apartment living, I would love a lap dog if possible but obviously the biggest concern is having them with my bird in the house. They would never EVER be left unsupervised or without one of them in a cage for safety, I would still like a dog that can live in peace with my bird. Some breeds I've heard were good were pomchi, tibetan spaniel, and cavaliers. At the same time I've heard they're the worst to house with a small animal..any help from people who have both? some insight to what you've had to do?","", "public", "",'2023-06-07 13:43:39.8098142+03:00', 0),
  (11, "d644ae49-0a31-4409-afea-cd8c5c3006de", "test", "Help me Identify this CREATURE!", "So, here is the story. I went to my shed today to get some gardening tools and I found a hole burrowed into my concrete flooring. I heard a weird growling noise coming from it and poked it with a stick. What came out was something TERRIFYING! It was a blend of a frog and a monster, I think it was actually Jabba the Hut but a smaller version of him. What should I do? He still lives in the hole!","", "public", "",'2022-07-03 23:55:11.8098142+03:00', 1),
  (2, "837fcac6-c188-49f3-8c1c-8f99dcfef5b3", "ForumJunkie", "Private user here", "Click on me I am a private user", "", "public", "", '2023-06-04 11:43:39.8098142+03:00', 0)
;

INSERT INTO comments (id, user_id, post_id, title, text, image, creation_date) VALUES
  (7, "12ef5526-049a-4857-b46d-fc90db00cb92", 6, "RE: I have a parrot but would like a lap dog too. What are some good breeds?", "I have both the pomchi and the tibetan spaniel as well as a parrot. Living together with them is absolute hell!", "", '2023-06-08 15:43:39.8098142+03:00'),
  (8, "945c7df5-88a2-4c58-bd6d-d92e04c6e282", 6, "RE: I have a parrot but would like a lap dog too. What are some good breeds?", "If you're looking for a dog that can handle apartment living and peacefully cohabitate with your bird, you might consider the 'Canine Zenitsu.' This breed has mastered the art of tranquility and meditation, making them the perfect companion for your bird's soothing chirps. They'll ensure a harmonious atmosphere in your home!", "", '2023-06-08 16:43:39.8098142+03:00'),
  (9, "12ef5526-049a-4857-b46d-fc90db00cb92", 6, "RE: I have a parrot but would like a lap dog too. What are some good breeds?", "UPDATE! I just got a cat instead. It was too hard to pick between different dogs.", "",'2023-06-08 19:22:31.8098142+03:00'),
  (10, "945c7df5-88a2-4c58-bd6d-d92e04c6e282", 3, "RE: Does your cats like listening to music?", "My cat also loves music. He especially loves the popular hit song Never Gonna Give You Up by Rick Astley.","",'2023-06-08 12:50:40.8098142+03:00')
;

INSERT INTO followers (follower, followed, allowed) VALUES 
  ("d644ae49-0a31-4409-afea-cd8c5c3006de","1c8e81c8-95d7-4d5a-a248-9d042f8e974c",0),
  ("d644ae49-0a31-4409-afea-cd8c5c3006de","de98c4ff-387c-474b-af2e-8dad191915fb",1),
  ("1c8e81c8-95d7-4d5a-a248-9d042f8e974c" , "d644ae49-0a31-4409-afea-cd8c5c3006de",0),
  ("de98c4ff-387c-474b-af2e-8dad191915fb","d644ae49-0a31-4409-afea-cd8c5c3006de",1),
  ("d644ae49-0a31-4409-afea-cd8c5c3006de", "945c7df5-88a2-4c58-bd6d-d92e04c6e282",true)
;