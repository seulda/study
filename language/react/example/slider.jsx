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
      <div className="carousel-content">
        <div className="carousel-content-left">
            <table className="table-style">
                <thead>
                    <tr>
                        <th>header</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>01</td>
                    </tr>
                    <tr>
                        <td>02</td>
                    </tr>
                    <tr>
                        <td>03</td>
                    </tr>
                </tbody>
            </table>
        </div>
        
        <div className="carousel-content-right">
            <Slider {...settings}>
                <div>
                    <table>
                        <thead>
                            <tr>
                                <th>1</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>Item1@@@@@@@@@@</td>
                            </tr>
                            <tr>
                                <td>Item2</td>
                            </tr>
                            <tr>
                                <td>Item3</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <div>
                    <table>
                        <thead>
                            <tr>
                                <th>2</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>Item1</td>
                            </tr>
                            <tr>
                                <td>Item2</td>
                            </tr>
                            <tr>
                                <td>Item3</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <div>
                    <table>
                        <thead>
                            <tr>
                                <th>3</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>Item1</td>
                            </tr>
                            <tr>
                                <td>Item2</td>
                            </tr>
                            <tr>
                                <td>Item3</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <div>
                    <table>
                        <thead>
                            <tr>
                                <th>4</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>Item1</td>
                            </tr>
                            <tr>
                                <td>Item2</td>
                            </tr>
                            <tr>
                                <td>Item3</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <div>
                    <table>
                        <thead>
                            <tr>
                                <th>5</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>Item1</td>
                            </tr>
                            <tr>
                                <td>Item2</td>
                            </tr>
                            <tr>
                                <td>Item3</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <div>
                    <table>
                        <thead>
                            <tr>
                                <th>6</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>Item1</td>
                            </tr>
                            <tr>
                                <td>Item2</td>
                            </tr>
                            <tr>
                                <td>Item3</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </Slider>
        </div>
      </div>
    );
  }
}


export default TableSlider;



// .slider-box {
//     // padding: 80px 30px;
//     margin: 5px;
//     border: 2px solid ; 
// }
// .carousel-content {
// 	display: flex; 
//     margin: 5px;
// }
// .carousel-content-left {
//     margin: 0px;
// 	padding: 0px;
// }
// .carousel-content-right {
//     width: 40vw;
// }
// .carousel-content-right table {
//     width: 100%;
//     color: white;
//     background-color: #6B6B6B;
//     text-align: center;
//     border: 2px solid;
//     border-left: none;
//     border-color: black;
// }
// .carousel-content-right th {
//     padding: 10px 0px;
//     border: 2px solid;
//     border-color: black;
// }
// .carousel-content-right td {
//     padding: 5px;
// }


// .table-style {
//     color: white;
//     background-color: #6B6B6B;
//     text-align: center;
//     border: 2px solid;
//     border-color: black;
// }
// .table-style th {
//     padding: 10px;
//     border: 2px solid;
//     border-color: black;
// }
// .table-style td {
//     padding: 5px;
// }
