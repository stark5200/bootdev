/* Blink Example

   This example code is in the Public Domain (or CC0 licensed, at your option.)

   Unless required by applicable law or agreed to in writing, this
   software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied.
*/
#include <stdio.h>
#include "freertos/FreeRTOS.h"
#include "freertos/task.h"
#include "driver/gpio.h"
#include "esp_log.h"
#include "led_strip.h"
#include "sdkconfig.h"

static const char *TAG = "example";

/* Use project configuration menu (idf.py menuconfig) to choose the GPIO to blink,
   or you can edit the following line and set a number here.
*/
#define BLINK_GPIO CONFIG_BLINK_GPIO

static uint8_t s_led_state = 0;
static uint8_t start_time = 0;
static uint8_t current_time = 0;
static uint8_t red = 0;
static uint8_t green = 0;
static uint8_t blue = 0;

#ifdef CONFIG_BLINK_LED_STRIP

static led_strip_handle_t led_strip;

// custom functions
void random_wurfel();
void egg_watch(int seconds, int r, int g, int b);
//int *decToBinary(int n);

static void blink_led(int a, int r, int g, int b)
{   
    /* If the addressable LED is enabled */
    if (s_led_state) {
        /* Set the LED pixel using RGB from 0 (0%) to 255 (100%) for each color */
        led_strip_set_pixel(led_strip, a, r, g, b);
        /* Refresh the strip to send data */
        led_strip_refresh(led_strip);
    } else {
        /* Set all LED off to clear all pixels */
        led_strip_clear(led_strip);
    }
}

static void configure_led(void)
{
    ESP_LOGI(TAG, "Example configured to blink addressable LED!");
    /* LED strip initialization with the GPIO and pixels number*/
    led_strip_config_t strip_config = {
        .strip_gpio_num = BLINK_GPIO,
        .max_leds = 25, // at least one LED on board
    };
#if CONFIG_BLINK_LED_STRIP_BACKEND_RMT
    led_strip_rmt_config_t rmt_config = {
        .resolution_hz = 10 * 1000 * 1000, // 10MHz
        .flags.with_dma = false,
    };
    ESP_ERROR_CHECK(led_strip_new_rmt_device(&strip_config, &rmt_config, &led_strip));
#elif CONFIG_BLINK_LED_STRIP_BACKEND_SPI
    led_strip_spi_config_t spi_config = {
        .spi_bus = SPI2_HOST,
        .flags.with_dma = true,
    };
    ESP_ERROR_CHECK(led_strip_new_spi_device(&strip_config, &spi_config, &led_strip));
#else
#error "unsupported LED strip backend"
#endif
    /* Set all LED off to clear all pixels */
    led_strip_clear(led_strip);
}

#elif CONFIG_BLINK_LED_GPIO

static void blink_led(void)
{
    /* Set the GPIO level according to the state (LOW or HIGH)*/
    gpio_set_level(BLINK_GPIO, s_led_state);
}

static void configure_led(void)
{
    ESP_LOGI(TAG, "Example configured to blink GPIO LED!");
    gpio_reset_pin(BLINK_GPIO);
    /* Set the GPIO as a push/pull output */
    gpio_set_direction(BLINK_GPIO, GPIO_MODE_OUTPUT);
}

#else
#error "unsupported LED type"
#endif

void app_main(void)
{

    /* Configure the peripheral according to the LED type */
    configure_led();

    //int x = 0;
    //int add = 0;
    s_led_state = 1;
    start_time = 125;
    current_time = start_time;
    red = 1;
    green = 50;
    blue = 1;
    // int btn_level = 0;

    while (1) {
        // choose function, random_würfel
        //random_wurfel();
        egg_watch(current_time, red, green, blue);
        if (current_time > 0) {
            current_time = current_time - 1;
        }
        red = (((start_time - current_time) * 50) / (start_time*2));
        green = ((current_time * 50) / start_time);

        // task delay
        vTaskDelay(CONFIG_BLINK_PERIOD / portTICK_PERIOD_MS);

        // clear led,  always execute last
        led_strip_clear(led_strip);
    }
}

void random_wurfel(void) {   
    int r = rand() % 6 + 1;
    int red = rand() % 30;
    int green = rand() % 30;
    int blue = rand() % 30;
    if (r == 1) {
        blink_led(12, red, green, blue);
    } else if (r == 2) {
        blink_led(4, red, green, blue);
        blink_led(20, red, green, blue);
    } else if (r == 3) {
        blink_led(0, red, green, blue);
        blink_led(12, red, green, blue);
        blink_led(24, red, green, blue);
    } else if (r == 4) {
        blink_led(0, red, green, blue);
        blink_led(4, red, green, blue);
        blink_led(20, red, green, blue);
        blink_led(24, red, green, blue);
    } else if (r == 5) {
        blink_led(0, red, green, blue);
        blink_led(4, red, green, blue);
        blink_led(12, red, green, blue);
        blink_led(20, red, green, blue);
        blink_led(24, red, green, blue);
    } else if (r == 6) {
        blink_led(0, red, green, blue);
        blink_led(4, red, green, blue);
        blink_led(10, red, green, blue);
        blink_led(14, red, green, blue);
        blink_led(20, red, green, blue);
        blink_led(24, red, green, blue);
    }

    ESP_LOGI(TAG, "you got a: %d", r);
};

void egg_watch(int seconds, int r, int g, int b) {
    if (seconds == 0) {
        int i = 0;
        while (i < 25) {
            blink_led(i, 50, 1, 1);
            i++;
        }
        ESP_LOGI(TAG, "Time is Up!!!");
        vTaskDelay(CONFIG_BLINK_PERIOD / portTICK_PERIOD_MS);
        // clear led,  always execute last
        led_strip_clear(led_strip);
    } else {
        int i = 0;
        while (i < 6) {
            if (i < 5) {
                if ((1 << i) & (seconds%60)) {
                    blink_led(4 - i, r, g, b);
                }
                if ((1 << i) & (seconds/60)) {
                    blink_led(19 - i, r, g, b);
                }
            } else {
                if ((1 << i) & (seconds%60)) {
                    blink_led(9 - (i-5), r, g, b);
                }
            }
            i++;
        }
    }
    /*
    uint8_t *m = decToBinary(seconds / 60);
    uint8_t *s = decToBinary(seconds % 60);
    if (m && s) {
        int i = 0;
        while (i < 6) {
            if (m[i] == 1) {
                blink_led(i, r, g, b);
            }
            if (s[i] == 1) {
                blink_led(15 + i, r, g, b);
            }
            i++;
        }
        free(m);
        free(s);
    }
    */
    ESP_LOGI(TAG, "Time remaining: %d seconds!", seconds);
};

/*
uint8_t *decToBinary(int n) {
    // array to store binary number
    uint8_t *binaryNum = calloc(6, sizeof(uint8_t));
    if (binaryNum == NULL) {
        return NULL;
    }
    // counter for binary array
    int i = 5;
    while (n >= 0) {
        // storing remainder in binary array
        binaryNum[i] = n % 2;
        n = n / 2;
        i = i - 1;
    }
    return binaryNum;
};
*/
