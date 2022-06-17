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
    export let path_api = "";
    let listgongjuday = [];
    let listsdsbnight = [];
    let size_image = "40";
    let size_clock = "50";
    let day_date_draw = "";
    let day_next_year = "";
    let day_next_month = "";
    let day_next_day = "";
    let day_prize1 = "";
    let day_prize2 = "";
    let day_prize3 = "";
    let day_1_prize1 = "0";
    let day_2_prize1 = "0";
    let day_3_prize1 = "0";
    let day_4_prize1 = "0";

    let day_img_1_prize1 = "images/bungapink-blank.png";
    let day_img_2_prize1 = "images/bungapeach-blank.png";
    let day_img_3_prize1 = "images/bungahijau-blank.png";
    let day_img_4_prize1 = "images/bungaungu-blank.png";
    let day_img_1_prize2 = "images/bolakecilbiru-blank.png";
    let day_img_2_prize2 = "images/bolakecilpink-blank.png";
    let day_img_3_prize2 = "images/bolakecilmerah-blank.png";
    let day_img_4_prize2 = "images/bolakecilungu-blank.png";
    let day_img_1_prize3 = "images/bolakecilbiru-blank.png";
    let day_img_2_prize3 = "images/bolakecilpink-blank.png";
    let day_img_3_prize3 = "images/bolakecilmerah-blank.png";
    let day_img_4_prize3 = "images/bolakecilungu-blank.png";
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
        day_next_year = dayjs(data["nextdraw"]).format("YYYY");
        day_next_month = dayjs(data["nextdraw"]).format("MM");
        day_next_day = dayjs(data["nextdraw"]).format("DD");
        day_prize1 = data["prize1"];
        day_prize2 = data["prize2"];
        day_prize3 = data["prize3"];
        day_1_prize1 = day_prize1[0];
        day_2_prize1 = day_prize1[1];
        day_3_prize1 = day_prize1[2];
        day_4_prize1 = day_prize1[3];
        // day_img_1_prize1 = getImage(day_prize1[0]);
        // day_img_2_prize1 = getImage(day_prize1[1]);
        // day_img_3_prize1 = getImage(day_prize1[2]);
        // day_img_4_prize1 = getImage(day_prize1[3]);
        // day_img_1_prize2 = getImage(day_prize2[0]);
        // day_img_2_prize2 = getImage(day_prize2[1]);
        // day_img_3_prize2 = getImage(day_prize2[2]);
        // day_img_4_prize2 = getImage(day_prize2[3]);
        // day_img_1_prize3 = getImage(day_prize3[0]);
        // day_img_2_prize3 = getImage(day_prize3[1]);
        // day_img_3_prize3 = getImage(day_prize3[2]);
        // day_img_4_prize3 = getImage(day_prize3[3]);
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
                urlimg = "images/bungaungu-0.png";
                break;
            case "1":
                urlimg = "images/bungaungu-1.png";
                break;
            case "2":
                urlimg = "images/bungaungu-1.png";
                break;
            case "3":
                urlimg = "images/bungaungu-1.png";
                break;
            case "4":
                urlimg = "images/bungaungu-1.png";
                break;
            case "5":
                urlimg = "images/bungaungu-1.png";
                break;
            case "6":
                urlimg = "images/bungaungu-1.png";
                break;
            case "7":
                urlimg = "images/bungaungu-1.png";
                break;
            case "8":
                urlimg = "images/bungaungu-1.png";
                break;
            case "9":
                urlimg = "images/bungaungu-1.png";
                break;
        }
        return urlimg;
    }
    async function initgongjuDAY() {
        const resPasar = await fetch(path_api+"api/listgongjuday", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({}),
        });
        if (!resPasar.ok) {
            const pasarMessage = `An error has occured: ${resPasar.status}`;
            throw new Error(pasarMessage);
        } else {
            const jsonPasar = await resPasar.json();
            if (jsonPasar.status == 200) {
                let record = jsonPasar.record;
                if (record != null) {
                    for (var i = 0; i < record.length; i++) {
                        listgongjuday = [
                            ...listgongjuday,
                            {
                                sdsbday_date: record[i]["sdsbday_date"],
                                sdsbday_prize1: record[i]["sdsbday_prize1"],
                                sdsbday_prize2: record[i]["sdsbday_prize2"],
                                sdsbday_prize3: record[i]["sdsbday_prize3"],
                            },
                        ];
                    }
                } else {
                    alert("Error");
                }
            } else {
                alert("Error");
            }
        }
    }
    // initgongjuDAY()
  </script>
<section class="hidden lg:flex gap-2 my-16">
    <Carousel />
</section>
<section class="my-16 w-full relative">
    <hr class="w-full bg-[pink] h-[2px] ">
    <h2 class="text-[pink] text-3xl text-center bg-white absolute -top-4 left-10 z-auto ">다음 그림</h2>
</section>
<section class="border-solid border-4  border-[#74aa63] my-16 rounded-3xl  p-5">
    <div class="flex flex-col my-2">
        <p class="text-[#74aa63] text-4xl font-poppins font-extrabold">공주의 날</p>
        <div class="w-full relative">
            <img class="w-full" src="images/day.png" alt="">
            <div class="flex justify-center  absolute top-8 w-full gap-36 -left-11 ">
                <span class="text-[150px] text-white text-center ">{temp_day_hour}</span>
                <span class="text-[150px] text-white text-center  ">{temp_day_minute}</span>
                <span class="text-[150px] text-white text-center ">{temp_day_second}</span>
            </div>
        </div>
    </div>
</section>
<section class="my-16 w-full relative">
    <hr class="w-full bg-[pink] h-[2px] ">
    <h2 class="text-[pink] text-3xl text-center bg-white absolute -top-4 left-10 z-auto">오늘의 출력</h2>
</section>
<section class="my-16 w-full relative">
    <img class="w-full z-0" src="images/background_xcferm.jpg" alt="">
    <div class="flex justify-items-center  w-full absolute top-10 left-10">
        <div class="text-white text-xs lg:text-2xl w-full mx-5">일등 상</div>
        <div class="text-white text-xs lg:text-2xl w-full text-right mr-20">{day_next_year}년 {day_next_month}월 {day_next_day}일</div>
    </div>
    <div class="flex justify-around   w-full absolute top-32 left-0 ">
        <img class="w-[200px]" src="{day_img_1_prize1}" alt="">
        <img class="w-[200px]" src="{day_img_2_prize1}" alt="">
        <img class="w-[200px]" src="{day_img_3_prize1}" alt="">
        <img class="w-[200px]" src="{day_img_4_prize1}" alt="">
    </div>
    <div class="flex justify-around w-full">
        <div class="flex flex-col absolute bottom-16 left-16 gap-2">
            <div class="text-white text-xs lg:text-2xl w-full">2등상</div>
            <div class="flex justify-start gap-20  w-full  ">
                <img class="" src="{day_img_1_prize2}" alt="">
                <img class="" src="{day_img_2_prize2}" alt="">
                <img class="" src="{day_img_3_prize2}" alt="">
                <img class="" src="{day_img_4_prize2}" alt="">
            </div>
        </div>
        <div class="flex flex-col absolute bottom-16 right-16 gap-2">
            <div class="text-white text-xs lg:text-2xl w-full">3등상</div>
            <div class="flex justify-start gap-20  w-full  ">
                <img class="" src="{day_img_1_prize3}" alt="">
                <img class="" src="{day_img_2_prize3}" alt="">
                <img class="" src="{day_img_3_prize3}" alt="">
                <img class="" src="{day_img_4_prize3}" alt="">
            </div>
        </div>
    </div>
    
</section>