/* Blink Example

   This example code is in the Public Domain (or CC0 licensed, at your option.)

   Unless required by applicable law or agreed to in writing, this
   software is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied.
*/
#include <stdio.h>
#include <stdlib.h>
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

/*
gpio_config_t gpioConfig = {
    .pin_bit_mask = (1 << LED1_GPIO) | (1 << LED2_GPIO),
    .mode = GPIO_MODE_OUTPUT,
    .pull_up_en = false,
    .pull_down_en = false,
    .intr_type = GPIO_INTR_DISABLE
};  
*/

// Trying to get button to work

//const gpio_num_t LED_PINS[] = { GPIO_NUM_27, GPIO_NUM_25, GPIO_NUM_32 };
/*
const uint8_t    LED_NUMBER = 3;

// defines the index of the active LED
uint8_t ledIndex = 0;

const uint16_t LOOP_FREQUENCY = 25;                    // Hz
const uint16_t WAIT_PERIOD    = 1000 / LOOP_FREQUENCY; // ms

typedef struct Timer {
    uint32_t laptime;
    uint32_t ticks;
} timer_t;

typedef struct Button {
    gpio_num_t pin;
    uint8_t    state;
} btn_t;

btn_t shiftButton = { GPIO_NUM_2, 0 }; // left button
btn_t sleepButton = { GPIO_NUM_9, 0 }; // right button
*/
/*
static void IRAM_ATTR button_isr_handler(void* arg) {
  int gpio_num = (int) arg;
  ESP_EARLY_LOGI(TAG, "Button pressed on GPIO %d", gpio_num);
}
*/


void app_main(void)
{   
    /* Configure the peripheral according to the LED type */
    configure_led();

    // setup for buttons
    //pinMode(shiftButton.pin, INPUT);
    //pinMode(sleepButton.pin, INPUT);

    // configures the LED pins
    /*
    for (uint8_t i=0; i<LED_NUMBER; i++) {
        pinMode(LED_PINS[i], OUTPUT);
    }
    */
    // turns on the active LED
    
    //digitalWrite(LED_PINS[ledIndex], HIGH);

    // initializes the timer
    //timer = { millis(), 0 };

    /*
    gpio_config_t button_config = {
        .pin_bit_mask = (1ULL << GPIO_NUM_0),
        .mode = GPIO_MODE_INPUT,
        .pull_up_en = GPIO_PULLUP_ENABLE,
        .intr_type = GPIO_INTR_NEGEDGE 
    };
    gpio_config(&button_config);
    gpio_install_isr_service(0);
    gpio_isr_handler_add(GPIO_NUM_0, button_isr_handler, (void*) GPIO_NUM_0);
    */


    //int x = 0;
    //int add = 0;
    s_led_state = 1;
    start_time = 119;
    current_time = start_time;
    red = 1;
    green = 50;
    blue = 1;
    // int btn_level = 0;
    while (1)
    {   
        /*
        if (gpio_get_level(BTN_GPIO) != btn_level) {
            btn_level = gpio_get_level(BTN_GPIO);
        }
        */
        
        //ESP_LOGI(TAG, "Turning the LED %s!", s_led_state == true ? "ON" : "OFF");
        //blink_led(r, red, green, blue);
        /* Toggle the LED state */
        //s_led_state = !s_led_state;
        //add = !add;
        //x = (x + 1) % 25;
        //if (add) {  x = (x + 1) % 25;}
        // vTaskDelay(1000 / 1);
        //led_strip_clear(led_strip);

        //button part
        /*
        readButton(&shiftButton);
        if (pressed(&shiftButton)) {
            ++ledIndex %= LED_NUMBER;
            updateLED();
        }
        waitForNextCycle();
        */

        // choose function, random_würfel
        //random_wurfel();
        egg_watch(current_time, red, green, blue);
        current_time = current_time - 1;
        red = (((start_time - current_time) * 50) / start_time);
        green = ((current_time * 50) / start_time);

        // task delay
        vTaskDelay(CONFIG_BLINK_PERIOD / portTICK_PERIOD_MS);

        // clear led,  always execute last
        led_strip_clear(led_strip);

    }
};

/*
void updateLED() {
    for (uint8_t i=0; i<LED_NUMBER; i++) {
        digitalWrite(LED_PINS[i], i == ledIndex ? HIGH : LOW);
    }
}
*/

// ----------------------------------------------------------------------------
// Button status handling
// ----------------------------------------------------------------------------
/*
void readButton(btn_t *b) {
    bool pressed = digitalRead(b->pin) == HIGH;

    if (pressed) {
             if (b->state < 0xfe) b->state++;
        else if (b->state == 0xfe) b->state = 2;
    } else if (b->state) {
        b->state = b->state == 0xff ? 0 : 0xff;
    }
}

bool pressed(btn_t *b) {
    return b->state == 1;
}

bool released(btn_t *b) {
    return b->state == 0xff;
}

bool held(btn_t *b) {
    return b->state > 1 && b->state < 0xff;
}
*/

// Time control of the main loop
// ----------------------------------------------------------------------------

/*
void waitForNextCycle() {
    uint32_t now;
    do { now = millis(); } while (now - timer.laptime < WAIT_PERIOD);
    timer.laptime = now;
    timer.ticks++;
}
*/

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
        //blink all red
        ESP_LOGI(TAG, "Time is Up!!!");
    } else {

        int i = 0;
        while (i < 6) {
            if ((1 << i) & (seconds/60)) {
                blink_led(i, r, g, b);
            }
            if ((1 << i) & (seconds%60)) {
                blink_led(15 + i, r, g, b);
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