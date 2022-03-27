import './App.css';
import { Autoplay, EffectFade, Virtual } from 'swiper';
import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/css';
import { useEffect, useState } from 'react';
import 'swiper/css/effect-fade';

function App() {
  const [photos, setPhotos] = useState([])

  useEffect(() => {
    const interval = setInterval(() => {
      fetch("/api/photos").then(res => res.json()).then((res) => {
        setPhotos(res)
      })
    }, 2000);
  
    return () => clearInterval(interval);
  }, []);

  return (

    <Swiper
      modules={[Virtual, EffectFade, Autoplay]}
      className='swiper'
      direction='horizontal'
      effect='fade'
      autoplay={{
        delay: 5000
      }}
      loop={true}
      // onSlideChange={() => console.log('slide change')}
      // onSwiper={(swiper) => console.log(swiper)}
    >
      <>
        {photos.map((photo, key) => {
          return (
            <SwiperSlide className="slide" key={key}>
              <img src={photo} rel='preload' />
            </SwiperSlide>
          )
          })}
      </>
    </Swiper>

  );
}

export default App;
