---
id: third-party-apis
sidebar_label: Third Party Api's
title: Third Party Api's
---

# Blackspace 3rd Party API's

## DataFiniti
DataFiniti provides us with a large catalog of business data from hundreds of online directories as well as websites. It's a really useful service since it 
currently witholds over 90 million unique business records with 99% accuracy across all records. Additionally more than 10 million businesses are updated on a weekly
basis, hence this is a fairly accurate tool for us to leverage.

Data is categorized by country and industry and also comes with reviews as well as geolocation. See below for a sample data point as well as potential query parameters

**business data schema - query**[https://developer.datafiniti.co/docs/business-data-schema]

**sample response**

```md
Name:	Staples
Location:	337 Hamilton St, Geneva, New York, US
Categories:	Copying and Duplicating Service, Office Equipment and Supplies, 
            Computer and Equipment Dealers, Geneva, Copiers & Supplies, Office Furniture and Equipment, 
            Furniture / Home Store, Printing Equipment & Supplies Dealers, Paper / Office Supplies Store, 
            Office Supplies Retail, Office Furniture & Equipment Dealers, Office Services, Supplies & Equipment, 
            Printing Services, Business & Professional Services, Computer Printers and Supplies, Print Shop, 
            Autographic Registers and Supplies, Filing Equipment Systems & Supplies
Phone #s:	2 phone(s) available
Emails:	1 email(s) available
Websites:	http://www.staples.com/, https://stores.staples.com/ny/geneva/337-hamilton-st, 
          https://stores.staples.com/ny/geneva/337-hamilton-st?cmpid=st_ll_gooandcid=LS:RETAIL:Y, 
          http://storelocator.staples.com/stores/en_US/NY/Geneva/1236
Avg Rating:	3.75 stars, 4 review(s) available
```

### Pricing
The free tier provides only 100 records per month and charges 10 dollars for 1000 additional records. Hence assuming we can properly manipulate
our query to return minority owned businesses, I anticipate with 100 dollars and 10000 company records spanning multiple industries would make out product pretty common-labels

**pricing**[https://datafiniti.co/pricing/business-data-pricing/]

## Yelp API
### Company Data
Yelp is also an interesting alternative providing us as a business with lots of business specific data. 

```json
{
  "id": "WavvLdfdP6g8aZTtbBQHTw",
  "alias": "gary-danko-san-francisco",
  "name": "Gary Danko",
  "image_url": "https://s3-media2.fl.yelpcdn.com/bphoto/CPc91bGzKBe95aM5edjhhQ/o.jpg",
  "is_claimed": true,
  "is_closed": false,
  "url": "https://www.yelp.com/biz/gary-danko-san-francisco?adjust_creative=wpr6gw4FnptTrk1CeT8POg&utm_campaign=yelp_api_v3&utm_medium=api_v3_business_lookup&utm_source=wpr6gw4FnptTrk1CeT8POg",
  "phone": "+14157492060",
  "display_phone": "(415) 749-2060",
  "review_count": 5296,
  "categories": [
    {
      "alias": "newamerican",
      "title": "American (New)"
    },
    {
      "alias": "french",
      "title": "French"
    },
    {
      "alias": "wine_bars",
      "title": "Wine Bars"
    }
  ],
  "rating": 4.5,
  "location": {
    "address1": "800 N Point St",
    "address2": "",
    "address3": "",
    "city": "San Francisco",
    "zip_code": "94109",
    "country": "US",
    "state": "CA",
    "display_address": [
      "800 N Point St",
      "San Francisco, CA 94109"
    ],
    "cross_streets": ""
  },
  "coordinates": {
    "latitude": 37.80587,
    "longitude": -122.42058
  },
  "photos": [
    "https://s3-media2.fl.yelpcdn.com/bphoto/CPc91bGzKBe95aM5edjhhQ/o.jpg",
    "https://s3-media4.fl.yelpcdn.com/bphoto/FmXn6cYO1Mm03UNO5cbOqw/o.jpg",
    "https://s3-media4.fl.yelpcdn.com/bphoto/HZVDyYaghwPl2kVbvHuHjA/o.jpg"
  ],
  "price": "$$$$",
  "hours": [
    {
      "open": [
        {
          "is_overnight": false,
          "start": "1730",
          "end": "2200",
          "day": 0
        },
        {
          "is_overnight": false,
          "start": "1730",
          "end": "2200",
          "day": 1
        },
        {
          "is_overnight": false,
          "start": "1730",
          "end": "2200",
          "day": 2
        },
        {
          "is_overnight": false,
          "start": "1730",
          "end": "2200",
          "day": 3
        },
        {
          "is_overnight": false,
          "start": "1730",
          "end": "2200",
          "day": 4
        },
        {
          "is_overnight": false,
          "start": "1730",
          "end": "2200",
          "day": 5
        },
        {
          "is_overnight": false,
          "start": "1730",
          "end": "2200",
          "day": 6
        }
      ],
      "hours_type": "REGULAR",
      "is_open_now": false
    }
  ],
  "transactions": [],
  "special_hours": [
    {
      "date": "2019-02-07",
      "is_closed": null,
      "start": "1600",
      "end": "2000",
      "is_overnight": false
    }
  ]
}
```

### Yelp search

One interesting takeaway is the fact that we could leverage YELP's apis (search) as people attempt to discover specific companies of interest instead of managing search from our end
Through yelp, we can allow our frontend users to search for companies based on attributes, categories, as well as a search term. The invocation of this api typically returns 1000 records
per query.

```json
{
  "total": 8228,
  "businesses": [
    {
      "rating": 4,
      "price": "$",
      "phone": "+14152520800",
      "id": "E8RJkjfdcwgtyoPMjQ_Olg",
      "alias": "four-barrel-coffee-san-francisco",
      "is_closed": false,
      "categories": [
        {
          "alias": "coffee",
          "title": "Coffee & Tea"
        }
      ],
      "review_count": 1738,
      "name": "Four Barrel Coffee",
      "url": "https://www.yelp.com/biz/four-barrel-coffee-san-francisco",
      "coordinates": {
        "latitude": 37.7670169511878,
        "longitude": -122.42184275
      },
      "image_url": "http://s3-media2.fl.yelpcdn.com/bphoto/MmgtASP3l_t4tPCL1iAsCg/o.jpg",
      "location": {
        "city": "San Francisco",
        "country": "US",
        "address2": "",
        "address3": "",
        "state": "CA",
        "address1": "375 Valencia St",
        "zip_code": "94103"
      },
      "distance": 1604.23,
      "transactions": ["pickup", "delivery"]
    },
    // ...
  ],
  "region": {
    "center": {
      "latitude": 37.767413217936834,
      "longitude": -122.42820739746094
    }
  }
}
```


#### Yelp business match
If we can consolidate a list of black owned businesses, we can invoke YELP's apis to gain more data about the business

```json
[
 {
   "name": "Gary Danko",
   "location": {
     "address1": "800 N Point St",
     "address2": "",
     "address3": "",
     "city": "San Francisco",
     "state": "CA",
     "zip_code": "94109",
     "country": "US"
   },
   "coordinates": {
     "latitude": 37.80587,
     "longitude": -122.42058
   },
   "phone": "+14157492060",
   "id": "WavvLdfdP6g8aZTtbBQHTw",
   "alias": "gary-danko-san-francisco"
 },
 {
   "name": "Good Grub Vending",
   "location": {
     "address1": "758 N Point St",
     "address2": "",
     "address3": "",
     "city": "San Francisco",
     "state": "CA",
     "zip_code": "94109",
     "country": "US"
   },
   "coordinates": {
     "latitude": 37.8061104,
     "longitude": -122.4195633
   },
   "phone": "+14157492060",
   "id": "fshpjHrtIlysFm0CnyUjbA",
   "alias": "good-grub-vending-san-francisco"
 }
]
```

## ZOMATO API 

**Zomato api link**[https://developers.zomato.com/documentation#!/restaurant/restaurant_0]

Super interesting api providing location specific collections of restaurants to the public.

#### Location Details
provides foodie index, nightlife index, top cuisines, and restaurants for a given Location
```json
{
  "popularity": {
    "popularity": "4.92",
    "nightlife_index": "4.95",
    "top_cuisines": [
      "cafe"
    ]
  },
  "location": {
    "entity_type": "group",
    "entity_id": "36932",
    "title": "Chelsea Market, Chelsea, New York City",
    "latitude": "40.742051",
    "longitude": "-74.004821",
    "city_id": "280",
    "city_name": "New York City",
    "country_id": "216",
    "country_name": "United States"
  },
  "best_rated_restaurants": [
    {
      "id": "16774318",
      "name": "Otto Enoteca & Pizzeria",
      "url": "https://www.zomato.com/new-york-city/otto-enoteca-pizzeria-greenwich-village",
      "location": {
        "address": "1 5th Avenue, New York, NY 10003",
        "locality": "Greenwich Village",
        "city": "New York City",
        "latitude": "40.732013",
        "longitude": "-73.996155",
        "zipcode": "10003",
        "country_id": "216"
      },
      "average_cost_for_two": "60",
      "price_range": "2",
      "currency": "$",
      "thumb": "https://b.zmtcdn.com/data/pictures/chains/8/16774318/a54deb9e4dbb79dd7c8091b30c642077_featured_thumb.png",
      "featured_image": "https://d.zmtcdn.com/data/pictures/chains/8/16774318/a54deb9e4dbb79dd7c8091b30c642077_featured_v2.png",
      "photos_url": "https://www.zomato.com/new-york-city/otto-enoteca-pizzeria-greenwich-village/photos#tabtop",
      "menu_url": "https://www.zomato.com/new-york-city/otto-enoteca-pizzeria-greenwich-village/menu#tabtop",
      "events_url": "https://www.zomato.com/new-york-city/otto-enoteca-pizzeria-greenwich-village/events#tabtop",
      "user_rating": {
        "aggregate_rating": "3.7",
        "rating_text": "Very Good",
        "rating_color": "5BA829",
        "votes": "1046"
      },
      "has_online_delivery": "0",
      "is_delivering_now": "0",
      "has_table_booking": "0",
      "deeplink": "zomato://r/16774318",
      "cuisines": "Cafe",
      "all_reviews_count": "15",
      "photo_count": "18",
      "phone_numbers": "(212) 228-2930",
      "photos": [
        {
          "id": "u_MjA5MjY1OTk5OT",
          "url": "https://b.zmtcdn.com/data/reviews_photos/c15/9eb13ceaf6e90129c276ce6ff980bc15_1435111695_640_640_thumb.JPG",
          "thumb_url": "https://b.zmtcdn.com/data/reviews_photos/c15/9eb13ceaf6e90129c276ce6ff980bc15_1435111695_200_thumb.JPG",
          "user": {
            "name": "John Doe",
            "zomato_handle": "John",
            "foodie_level": "Super Foodie",
            "foodie_level_num": "9",
            "foodie_color": "f58552",
            "profile_url": "https://www.zomato.com/john",
            "profile_deeplink": "zoma.to/u/1170245",
            "profile_image": "string"
          },
          "res_id": "16782899",
          "caption": "#awesome",
          "timestamp": "1435111770",
          "friendly_time": "3 months ago",
          "width": "640",
          "height": "640",
          "comments_count": "0",
          "likes_count": "0"
        }
      ],
      "all_reviews": [
        {
          "rating": "5",
          "review_text": "The best latte I've ever had. It tasted a little sweet",
          "id": "24127336",
          "rating_color": "305D02",
          "review_time_friendly": "2 months ago",
          "rating_text": "Insane!",
          "timestamp": "1435507367",
          "likes": "0",
          "user": {
            "name": "John Doe",
            "zomato_handle": "John",
            "foodie_level": "Super Foodie",
            "foodie_level_num": "9",
            "foodie_color": "f58552",
            "profile_url": "https://www.zomato.com/john",
            "profile_deeplink": "zoma.to/u/1170245",
            "profile_image": "string"
          },
          "comments_count": "0"
        }
      ]
    }
  ]
}
```

### Restaurant Details

Provides detailed restaurant menu
```json
{
  "daily_menu": [
    {
      "daily_menu_id": "16507624",
      "name": "Vinohradský pivovar",
      "start_date": "2016-03-08 11:00",
      "end_date": "2016-03-08 15:00",
      "dishes": [
        {
          "dish_id": "104089345",
          "name": "Tatarák ze sumce s toustem",
          "price": "149 Kč"
        }
      ]
    }
  ]
}
```

### Restaurants 

Provides actual restaurant details by id  
```json
{
  "id": "16774318",
  "name": "Otto Enoteca & Pizzeria",
  "url": "https://www.zomato.com/new-york-city/otto-enoteca-pizzeria-greenwich-village",
  "location": {
    "address": "1 5th Avenue, New York, NY 10003",
    "locality": "Greenwich Village",
    "city": "New York City",
    "latitude": "40.732013",
    "longitude": "-73.996155",
    "zipcode": "10003",
    "country_id": "216"
  },
  "average_cost_for_two": "60",
  "price_range": "2",
  "currency": "$",
  "thumb": "https://b.zmtcdn.com/data/pictures/chains/8/16774318/a54deb9e4dbb79dd7c8091b30c642077_featured_thumb.png",
  "featured_image": "https://d.zmtcdn.com/data/pictures/chains/8/16774318/a54deb9e4dbb79dd7c8091b30c642077_featured_v2.png",
  "photos_url": "https://www.zomato.com/new-york-city/otto-enoteca-pizzeria-greenwich-village/photos#tabtop",
  "menu_url": "https://www.zomato.com/new-york-city/otto-enoteca-pizzeria-greenwich-village/menu#tabtop",
  "events_url": "https://www.zomato.com/new-york-city/otto-enoteca-pizzeria-greenwich-village/events#tabtop",
  "user_rating": {
    "aggregate_rating": "3.7",
    "rating_text": "Very Good",
    "rating_color": "5BA829",
    "votes": "1046"
  },
  "has_online_delivery": "0",
  "is_delivering_now": "0",
  "has_table_booking": "0",
  "deeplink": "zomato://r/16774318",
  "cuisines": "Cafe",
  "all_reviews_count": "15",
  "photo_count": "18",
  "phone_numbers": "(212) 228-2930",
  "photos": [
    {
      "id": "u_MjA5MjY1OTk5OT",
      "url": "https://b.zmtcdn.com/data/reviews_photos/c15/9eb13ceaf6e90129c276ce6ff980bc15_1435111695_640_640_thumb.JPG",
      "thumb_url": "https://b.zmtcdn.com/data/reviews_photos/c15/9eb13ceaf6e90129c276ce6ff980bc15_1435111695_200_thumb.JPG",
      "user": {
        "name": "John Doe",
        "zomato_handle": "John",
        "foodie_level": "Super Foodie",
        "foodie_level_num": "9",
        "foodie_color": "f58552",
        "profile_url": "https://www.zomato.com/john",
        "profile_deeplink": "zoma.to/u/1170245",
        "profile_image": "string"
      },
      "res_id": "16782899",
      "caption": "#awesome",
      "timestamp": "1435111770",
      "friendly_time": "3 months ago",
      "width": "640",
      "height": "640",
      "comments_count": "0",
      "likes_count": "0"
    }
  ],
  "all_reviews": [
    {
      "rating": "5",
      "review_text": "The best latte I've ever had. It tasted a little sweet",
      "id": "24127336",
      "rating_color": "305D02",
      "review_time_friendly": "2 months ago",
      "rating_text": "Insane!",
      "timestamp": "1435507367",
      "likes": "0",
      "user": {
        "name": "John Doe",
        "zomato_handle": "John",
        "foodie_level": "Super Foodie",
        "foodie_level_num": "9",
        "foodie_color": "f58552",
        "profile_url": "https://www.zomato.com/john",
        "profile_deeplink": "zoma.to/u/1170245",
        "profile_image": "string"
      },
      "comments_count": "0"
    }
  ]
}
```

### Zomato search
Provides search functionality by cuisine, establishment, collection and location

```json
{
  "results_found": "53",
  "results_start": "11",
  "results_shown": "10",
  "restaurants": [
    {
      "id": "16774318",
      "name": "Otto Enoteca & Pizzeria",
      "url": "https://www.zomato.com/new-york-city/otto-enoteca-pizzeria-greenwich-village",
      "location": {
        "address": "1 5th Avenue, New York, NY 10003",
        "locality": "Greenwich Village",
        "city": "New York City",
        "latitude": "40.732013",
        "longitude": "-73.996155",
        "zipcode": "10003",
        "country_id": "216"
      },
      "average_cost_for_two": "60",
      "price_range": "2",
      "currency": "$",
      "thumb": "https://b.zmtcdn.com/data/pictures/chains/8/16774318/a54deb9e4dbb79dd7c8091b30c642077_featured_thumb.png",
      "featured_image": "https://d.zmtcdn.com/data/pictures/chains/8/16774318/a54deb9e4dbb79dd7c8091b30c642077_featured_v2.png",
      "photos_url": "https://www.zomato.com/new-york-city/otto-enoteca-pizzeria-greenwich-village/photos#tabtop",
      "menu_url": "https://www.zomato.com/new-york-city/otto-enoteca-pizzeria-greenwich-village/menu#tabtop",
      "events_url": "https://www.zomato.com/new-york-city/otto-enoteca-pizzeria-greenwich-village/events#tabtop",
      "user_rating": {
        "aggregate_rating": "3.7",
        "rating_text": "Very Good",
        "rating_color": "5BA829",
        "votes": "1046"
      },
      "has_online_delivery": "0",
      "is_delivering_now": "0",
      "has_table_booking": "0",
      "deeplink": "zomato://r/16774318",
      "cuisines": "Cafe",
      "all_reviews_count": "15",
      "photo_count": "18",
      "phone_numbers": "(212) 228-2930",
      "photos": [
        {
          "id": "u_MjA5MjY1OTk5OT",
          "url": "https://b.zmtcdn.com/data/reviews_photos/c15/9eb13ceaf6e90129c276ce6ff980bc15_1435111695_640_640_thumb.JPG",
          "thumb_url": "https://b.zmtcdn.com/data/reviews_photos/c15/9eb13ceaf6e90129c276ce6ff980bc15_1435111695_200_thumb.JPG",
          "user": {
            "name": "John Doe",
            "zomato_handle": "John",
            "foodie_level": "Super Foodie",
            "foodie_level_num": "9",
            "foodie_color": "f58552",
            "profile_url": "https://www.zomato.com/john",
            "profile_deeplink": "zoma.to/u/1170245",
            "profile_image": "string"
          },
          "res_id": "16782899",
          "caption": "#awesome",
          "timestamp": "1435111770",
          "friendly_time": "3 months ago",
          "width": "640",
          "height": "640",
          "comments_count": "0",
          "likes_count": "0"
        }
      ],
      "all_reviews": [
        {
          "rating": "5",
          "review_text": "The best latte I've ever had. It tasted a little sweet",
          "id": "24127336",
          "rating_color": "305D02",
          "review_time_friendly": "2 months ago",
          "rating_text": "Insane!",
          "timestamp": "1435507367",
          "likes": "0",
          "user": {
            "name": "John Doe",
            "zomato_handle": "John",
            "foodie_level": "Super Foodie",
            "foodie_level_num": "9",
            "foodie_color": "f58552",
            "profile_url": "https://www.zomato.com/john",
            "profile_deeplink": "zoma.to/u/1170245",
            "profile_image": "string"
          },
          "comments_count": "0"
        }
      ]
    }
  ]
}
```

