import React, { Component } from "react";
import Slider from "react-slick";

import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";

// npm install react-slick
// npm install slick-carousel


class TableSlider extends Component {
  render() {
    const settings = {
      dots: false,
      arrows: false,
      infinite: true,
      speed: 500,
      slidesToShow: 3,
      slidesToScroll: 1,
      autoplay: true,
      autoplaySpeed: 1500,
      pauseOnHover: false,
    };
    return (
      <div>
        <h2>table</h2>
        <table>
            <thead>
                <tr>
                    <th>A_header</th>
                    <th>B_header</th>
                    <th>C_header</th>
                    <th>D_header</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>A-1</td>
                    <td>B-1</td>
                    <td>C-1</td>
                    <td>D-1</td>
                </tr>
                <tr>
                    <td>A-2</td>
                    <td>B-2</td>
                    <td>C-2</td>
                    <td>D-2</td>
                </tr>
                <tr>
                    <td>A-3</td>
                    <td>B-3</td>
                    <td>C-3</td>
                    <td>D-3</td>
                </tr>
            </tbody>
        </table>
        <h2>slider</h2>
        <Slider {...settings}>
          <div>
            <h3>1</h3>
          </div>
          <div>
            <h3>2</h3>
          </div>
          <div>
            <h3>3</h3>
          </div>
          <div>
            <h3>4</h3>
          </div>
          <div>
            <h3>5</h3>
          </div>
          <div>
            <h3>6</h3>
          </div>
        </Slider>
      </div>
    );
  }
}


export default TableSlider;
