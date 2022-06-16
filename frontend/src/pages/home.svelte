<script>
    import { initializeApp } from "firebase/app";
    import { getDatabase, ref, onValue } from "firebase/database";
    import dayjs from "dayjs";
    import Carousel from '../lib/Carousel.svelte'

    const firebaseConfig = {
        apiKey: "AIzaSyBqVRbGvJBb1JEfKYyN6jgocZjzsx2lN2A",
        authDomain: "dazzling-pillar-328210.firebaseapp.com",
        databaseURL: "https://dazzling-pillar-328210-default-rtdb.asia-southeast1.firebasedatabase.app",
        projectId: "dazzling-pillar-328210",
        storageBucket: "dazzling-pillar-328210.appspot.com",
        messagingSenderId: "770359422621",
        appId: "1:770359422621:web:7933922e00547dc735ee74"
    };
    const app = initializeApp(firebaseConfig);
    const db = getDatabase(app);
    const sdsb4dday = ref(db, "sdsb4dday");
    const sdsb4dnight = ref(db, "sdsb4dnight");
    let listsdsbday = [];
    let listsdsbnight = [];
    let size_image = "40";
    let size_clock = "50";
    let day_date_draw = "";
    let day_next_draw = "";
    let day_prize1 = "";
    let day_prize2 = "";
    let day_prize3 = "";
    let day_img_1_prize1 = "number/ball-null.svg";
    let day_img_2_prize1 = "number/ball-null.svg";
    let day_img_3_prize1 = "number/ball-null.svg";
    let day_img_4_prize1 = "number/ball-null.svg";
    let day_img_1_prize2 = "number/ball-null.svg";
    let day_img_2_prize2 = "number/ball-null.svg";
    let day_img_3_prize2 = "number/ball-null.svg";
    let day_img_4_prize2 = "number/ball-null.svg";
    let day_img_1_prize3 = "number/ball-null.svg";
    let day_img_2_prize3 = "number/ball-null.svg";
    let day_img_3_prize3 = "number/ball-null.svg";
    let day_img_4_prize3 = "number/ball-null.svg";
    let night_date_draw = "";
    let night_next_draw = "";
    let night_prize1 = "";
    let night_prize2 = "";
    let night_prize3 = "";
    let night_img_1_prize1 = "number/ball-null.svg";
    let night_img_2_prize1 = "number/ball-null.svg";
    let night_img_3_prize1 = "number/ball-null.svg";
    let night_img_4_prize1 = "number/ball-null.svg";
    let night_img_1_prize2 = "number/ball-null.svg";
    let night_img_2_prize2 = "number/ball-null.svg";
    let night_img_3_prize2 = "number/ball-null.svg";
    let night_img_4_prize2 = "number/ball-null.svg";
    let night_img_1_prize3 = "number/ball-null.svg";
    let night_img_2_prize3 = "number/ball-null.svg";
    let night_img_3_prize3 = "number/ball-null.svg";
    let night_img_4_prize3 = "number/ball-null.svg";
    let temp_day = "";
    let temp_day_hour = "00";
    let temp_day_minute = "00";
    let temp_day_second = "00";
    onValue(sdsb4dday, (snapshot) => {
        const data = snapshot.val();
        day_date_draw = data["datedraw"];
        day_next_draw = dayjs(data["nextdraw"]).format("DD-MMM-YYYY");
        day_prize1 = data["prize1"];
        day_prize2 = data["prize2"];
        day_prize3 = data["prize3"];
        console.log(day_prize1)
        day_img_1_prize1 = getImage(day_prize1[0]);
        day_img_2_prize1 = getImage(day_prize1[1]);
        day_img_3_prize1 = getImage(day_prize1[2]);
        day_img_4_prize1 = getImage(day_prize1[3]);
        day_img_1_prize2 = getImage(day_prize2[0]);
        day_img_2_prize2 = getImage(day_prize2[1]);
        day_img_3_prize2 = getImage(day_prize2[2]);
        day_img_4_prize2 = getImage(day_prize2[3]);
        day_img_1_prize3 = getImage(day_prize3[0]);
        day_img_2_prize3 = getImage(day_prize3[1]);
        day_img_3_prize3 = getImage(day_prize3[2]);
        day_img_4_prize3 = getImage(day_prize3[3]);
        temp_day = dayjs(data["nextdraw"] + " 15:00:00").valueOf();
        setInterval(function () {
            let now = new Date().getTime();
            let distance = temp_day - now;
            let hours = Math.floor(
                (distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60)
            );
            let minutes = Math.floor(
                (distance % (1000 * 60 * 60)) / (1000 * 60)
            );
            let seconds = Math.floor((distance % (1000 * 60)) / 1000);
            temp_day_hour = hours;
            temp_day_minute = minutes;
            temp_day_second = seconds;
            if (hours < 10) {
                temp_day_hour = "0" + hours;
            } else {
                temp_day_hour = hours;
            }
            if (minutes < 10) {
                temp_day_minute = "0" + minutes;
            } else {
                temp_day_minute = minutes;
            }
            if (seconds < 10) {
                temp_day_second = "0" + seconds;
            } else {
                temp_day_second = seconds;
            }
            if (distance < 0) {
                clearInterval();
                temp_day_hour = "00";
                temp_day_minute = "00";
                temp_day_second = "00";
            }
            
        }, 1000);
    });
    let temp_night = "";
    let temp_night_hour = "00";
    let temp_night_minute = "00";
    let temp_night_second = "00";
    onValue(sdsb4dnight, (snapshot) => {
        const data = snapshot.val();
        night_date_draw = data["datedraw"];
        night_next_draw = dayjs(data["nextdraw"]).format("DD-MMM-YYYY");
        night_prize1 = data["prize1"];
        night_prize2 = data["prize2"];
        night_prize3 = data["prize3"];
        night_img_1_prize1 = getImage(night_prize1[0]);
        night_img_2_prize1 = getImage(night_prize1[1]);
        night_img_3_prize1 = getImage(night_prize1[2]);
        night_img_4_prize1 = getImage(night_prize1[3]);
        night_img_1_prize2 = getImage(night_prize2[0]);
        night_img_2_prize2 = getImage(night_prize2[1]);
        night_img_3_prize2 = getImage(night_prize2[2]);
        night_img_4_prize2 = getImage(night_prize2[3]);
        night_img_1_prize3 = getImage(night_prize3[0]);
        night_img_2_prize3 = getImage(night_prize3[1]);
        night_img_3_prize3 = getImage(night_prize3[2]);
        night_img_4_prize3 = getImage(night_prize3[3]);
        temp_night = dayjs(data["nextdraw"] + " 21:00:00").valueOf();
        setInterval(function () {
            let now = new Date().getTime();
            let distance = temp_night - now;
            let hours = Math.floor(
                (distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60)
            );
            let minutes = Math.floor(
                (distance % (1000 * 60 * 60)) / (1000 * 60)
            );
            let seconds = Math.floor((distance % (1000 * 60)) / 1000);
            if (hours < 10) {
                temp_night_hour = "0" + hours;
            } else {
                temp_night_hour = hours;
            }
            if (minutes < 10) {
                temp_night_minute = "0" + minutes;
            } else {
                temp_night_minute = minutes;
            }
            if (seconds < 10) {
                temp_night_second = "0" + seconds;
            } else {
                temp_night_second = seconds;
            }
            if (distance < 0) {
                clearInterval();
                temp_night_hour = "00";
                temp_night_minute = "00";
                temp_night_second = "00";
            }
        }, 1000);
    });
    function getImage(e) {
        let urlimg = "";
        switch (e) {
            case "0":
                urlimg = "number/ball-0.svg";
                break;
            case "1":
                urlimg = "number/ball-1.svg";
                break;
            case "2":
                urlimg = "number/ball-2.svg";
                break;
            case "3":
                urlimg = "number/ball-3.svg";
                break;
            case "4":
                urlimg = "number/ball-4.svg";
                break;
            case "5":
                urlimg = "number/ball-5.svg";
                break;
            case "6":
                urlimg = "number/ball-6.svg";
                break;
            case "7":
                urlimg = "number/ball-7.svg";
                break;
            case "8":
                urlimg = "number/ball-8.svg";
                break;
            case "9":
                urlimg = "number/ball-9.svg";
                break;
        }
        return urlimg;
    }
  </script>
<section class="flex gap-2 my-16">
    <Carousel />
</section>
<section class="my-16 w-full relative">
    <hr class="w-full bg-[pink] h-[2px] ">
    <h2 class="text-[pink] text-3xl text-center bg-white absolute -top-3 left-10 z-auto w-1/6">다음 그림</h2>
</section>
<section class="border-solid border-4  border-[#74aa63] my-16 rounded-lg  p-5">
    <div class="flex flex-col my-16">
        <p class="text-[#74aa63] text-4xl font-poppins font-extrabold">아침(왼쪽)과 밤(오른쪽) 출애굽</p>
        <div class="grid grid-cols-2 gap-2 w-full my-16">
            <div class="w-full relative">
                <img src="images/night.png" alt="">
                <div class="flex justify-center gap-16 absolute top-10 left-16">
                <span class="text-6xl text-white">{temp_day_hour}</span>
                <span class="text-6xl text-white mx-4">{temp_day_minute}</span>
                <span class="text-6xl text-white -mx-2">{temp_day_second}</span>
                </div>
            </div>
            <div class="w-full relative">
                <img src="images/day.png" alt="">
                <div class="flex justify-center gap-16 absolute top-10 left-20">
                <span class="text-6xl text-white">{temp_night_hour}</span>
                <span class="text-6xl text-white mx-4">{temp_night_minute}</span>
                <span class="text-6xl text-white -mx-2">{temp_night_second}</span>
                </div>
            </div>
        </div>
    </div>
</section>
<section class="my-16 w-full relative">
    <hr class="w-full bg-[pink] h-[2px] ">
    <h2 class="text-[pink] text-3xl text-center bg-white absolute -top-3 left-10 z-auto w-1/6">오늘의 출력</h2>
</section>
<section class="my-16 w-full relative">
    <img class="w-full" src="images/keluaran.jpg" alt="">
</section>