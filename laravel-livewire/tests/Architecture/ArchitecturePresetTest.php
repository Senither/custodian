<?php

pest()->project()->github('senither/custodian');

arch()->preset()->laravel();
arch()->preset()->security();
arch()->preset()->relaxed();
