import './App.css';
import { Autoplay, EffectCards, EffectCoverflow, EffectCreative, EffectCube, EffectFade, EffectFlip } from 'swiper';
import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/css';
import { useEffect, useState } from 'react';
import 'swiper/css/effect-fade';
import "swiper/css/effect-cards";
import "swiper/css/effect-coverflow";
import "swiper/css/effect-creative";
import "swiper/css/effect-cube";
import "swiper/css/effect-flip";

function App() {
  const [photos, setPhotos] = useState([])
  const [settings, setSettings] = useState({
    RefreshTime: 30000,
    PhotoTime: 10000
  })

  useEffect(() => {
    fetch('/api/settings').then(res => res.json()).then((res) => {
      setSettings(res)
    })
  }, [])

  useEffect(() => {
    const fetchPhotos = () => {
      const { innerWidth: width, innerHeight: height } = window;

      fetch(`/api/photos?w=${width}&h=${height}`).then(res => res.json()).then((res) => {
        setPhotos(res)
      })
    }

    const interval = setInterval(fetchPhotos, settings.RefreshTime);

    fetchPhotos()
  
    return () => clearInterval(interval);
  }, [settings]);

  return (

    <Swiper
      modules={[Autoplay, EffectFade]}
      className='swiper'
      direction='horizontal'
      effect='fade'
      autoplay={{
        delay: settings.PhotoTime
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
