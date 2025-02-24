package main

import (
	_ "myproject/www"        // ✅ Import for init() execution only
	_ "myproject/www/v1/sum" // ✅ Ensure sum init() runs
)




