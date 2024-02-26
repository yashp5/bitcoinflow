<template>
  <div>
    <apexchart
      :width="'100%'"
      :height="'250'"
      type="donut"
      :options="options"
      :series="series"
    ></apexchart>
  </div>
</template>

<script>
export default {
  props: {
    series: Array, // Use series passed from parent component
  },
  data: function () {
    return {
      options: {
        chart: {
          type: "donut",
          background: "transparent", // Enhances the glassmorphism effect
          toolbar: {
            show: false, // Hides the toolbar for a cleaner look
          },
        },
        theme: {
          mode: "dark", // Ensures compatibility with dark themes
        },
        labels: ["Put Volume", "Call Volume"],
        legend: {
          position: "bottom",
          labels: {
            colors: ["#f3f4f6"], // Light colors for visibility on dark background
          },
          fontFamily: "sans-serif", // Customize as per your UI design
          fontSize: "14px",
          markers: {
            fillColors: ["#f3f4f6", "#f3f4f6"], // Marker colors can also be adjusted
          },
        },
        plotOptions: {
          pie: {
            donut: {
              labels: {
                show: true,
                name: {
                  show: true,
                },
                // value: {
                //   show: false, // Disables the value labels
                // },
                // total: {
                //   show: false, // Disables the total value label
                // },
              },
            },
          },
        },
        stroke: {
          show: false, // Removes the border around the donut slices
        },
        fill: {
          type: "gradient", // Adds a gradient for a more dynamic look
        },
        dataLabels: {
          enabled: false, // Disables data labels to remove percentages
        },
        // Additional styling for the dark glass theme
        states: {
          hover: {
            filter: {
              type: "darken",
              value: 0.75,
            },
          },
        },
        responsive: [
          {
            breakpoint: 480,
            options: {
              chart: {
                width: "100%",
              },
              legend: {
                position: "bottom",
              },
            },
          },
        ],
      },
    };
  },
};
</script>
