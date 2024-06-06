import React, { useEffect, useRef, useState } from "react";
import Box from "@mui/material/Box";
import IconButton from "@mui/material/IconButton";
import ArrowBackIosIcon from "@mui/icons-material/ArrowBackIos";
import ArrowForwardIosIcon from "@mui/icons-material/ArrowForwardIos";
import { styled } from "@mui/system";

const CarouselContainer = styled(Box)({
  display: "flex",
  overflow: "hidden",
  width: "100%",
  position: "relative",
});

const CarouselItem = styled(Box)({
  minWidth: "100%",
  transition: "transform 0.5s ease-in-out",
});

const Carousel: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [index, setIndex] = useState(0);
  const totalItems = React.Children.count(children);
  const intervalRef = useRef<NodeJS.Timeout | null>(null);

  useEffect(() => {
    startAutoScroll();
    return () => stopAutoScroll();
  }, [index]);

  const startAutoScroll = () => {
    stopAutoScroll(); // Clear any existing intervals
    intervalRef.current = setInterval(() => {
      setIndex((prevIndex) => (prevIndex + 1) % totalItems);
    }, 3000); // Change every 3 seconds
  };

  const stopAutoScroll = () => {
    if (intervalRef.current) {
      clearInterval(intervalRef.current);
      intervalRef.current = null;
    }
  };

  const next = () => {
    setIndex((prevIndex) => (prevIndex + 1) % totalItems);
  };

  const prev = () => {
    setIndex((prevIndex) => (prevIndex - 1 + totalItems) % totalItems);
  };

  return (
    <Box sx={{ position: "relative" }}>
      <CarouselContainer>
        <Box
          sx={{
            display: "flex",
            transform: `translateX(-${index * 100}%)`,
            transition: "transform 0.5s ease-in-out",
          }}
        >
          {React.Children.map(children, (child, idx) => (
            <CarouselItem key={idx}>{child}</CarouselItem>
          ))}
        </Box>
      </CarouselContainer>
      <IconButton
        onClick={prev}
        sx={{
          position: "absolute",
          top: "50%",
          left: "0",
          transform: "translateY(-50%)",
        }}
      >
        <ArrowBackIosIcon />
      </IconButton>
      <IconButton
        onClick={next}
        sx={{
          position: "absolute",
          top: "50%",
          right: "0",
          transform: "translateY(-50%)",
        }}
      >
        <ArrowForwardIosIcon />
      </IconButton>
    </Box>
  );
};

export default Carousel;
