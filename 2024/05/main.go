package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//input := "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13\n\n75,47,61,53,29\n97,61,53,29,13\n75,29,13\n75,97,47,61,53\n61,13,29\n97,13,75,29,47"
	input := "56|54\n94|55\n94|11\n42|41\n42|21\n42|32\n54|13\n54|42\n54|61\n54|98\n26|95\n26|79\n26|11\n26|14\n26|24\n38|41\n38|74\n38|54\n38|32\n38|39\n38|94\n16|37\n16|38\n16|55\n16|14\n16|29\n16|44\n16|24\n95|88\n95|62\n95|32\n95|34\n95|11\n95|91\n95|57\n95|65\n57|42\n57|39\n57|14\n57|65\n57|17\n57|61\n57|19\n57|21\n57|64\n76|97\n76|57\n76|95\n76|15\n76|44\n76|54\n76|11\n76|61\n76|64\n76|91\n35|58\n35|94\n35|32\n35|79\n35|91\n35|16\n35|39\n35|57\n35|62\n35|21\n35|17\n98|38\n98|95\n98|35\n98|76\n98|56\n98|23\n98|62\n98|97\n98|55\n98|67\n98|88\n98|33\n62|17\n62|15\n62|16\n62|32\n62|58\n62|65\n62|21\n62|61\n62|54\n62|26\n62|73\n62|79\n62|74\n84|73\n84|75\n84|74\n84|44\n84|27\n84|15\n84|76\n84|34\n84|37\n84|88\n84|35\n84|62\n84|91\n84|33\n65|41\n65|24\n65|58\n65|79\n65|26\n65|13\n65|23\n65|87\n65|19\n65|67\n65|14\n65|29\n65|98\n65|39\n65|21\n88|57\n88|54\n88|62\n88|38\n88|21\n88|17\n88|37\n88|46\n88|39\n88|73\n88|42\n88|25\n88|19\n88|74\n88|61\n88|65\n29|97\n29|11\n29|91\n29|35\n29|84\n29|57\n29|75\n29|76\n29|38\n29|34\n29|15\n29|37\n29|33\n29|44\n29|56\n29|27\n29|88\n15|24\n15|16\n15|26\n15|42\n15|19\n15|64\n15|14\n15|28\n15|58\n15|39\n15|17\n15|67\n15|61\n15|65\n15|23\n15|13\n15|32\n15|41\n19|87\n19|21\n19|39\n19|23\n19|76\n19|13\n19|34\n19|24\n19|79\n19|58\n19|14\n19|75\n19|95\n19|29\n19|41\n19|44\n19|98\n19|94\n19|16\n58|27\n58|87\n58|84\n58|11\n58|55\n58|97\n58|23\n58|67\n58|98\n58|79\n58|34\n58|95\n58|76\n58|56\n58|13\n58|28\n58|94\n58|75\n58|24\n58|14\n17|84\n17|76\n17|19\n17|23\n17|75\n17|28\n17|29\n17|58\n17|94\n17|13\n17|41\n17|16\n17|87\n17|24\n17|67\n17|14\n17|21\n17|55\n17|79\n17|39\n17|26\n33|39\n33|19\n33|54\n33|62\n33|94\n33|58\n33|13\n33|35\n33|61\n33|21\n33|26\n33|41\n33|57\n33|65\n33|17\n33|64\n33|42\n33|25\n33|79\n33|74\n33|73\n33|91\n11|32\n11|62\n11|57\n11|46\n11|35\n11|74\n11|38\n11|39\n11|61\n11|33\n11|42\n11|17\n11|15\n11|19\n11|41\n11|64\n11|21\n11|54\n11|91\n11|73\n11|25\n11|65\n11|88\n44|56\n44|32\n44|88\n44|38\n44|11\n44|91\n44|35\n44|74\n44|54\n44|95\n44|46\n44|57\n44|97\n44|65\n44|33\n44|42\n44|27\n44|61\n44|73\n44|15\n44|37\n44|62\n44|64\n44|34\n39|67\n39|94\n39|34\n39|26\n39|97\n39|76\n39|44\n39|58\n39|79\n39|24\n39|87\n39|95\n39|27\n39|55\n39|29\n39|16\n39|84\n39|28\n39|56\n39|23\n39|75\n39|98\n39|13\n39|14\n91|28\n91|73\n91|16\n91|58\n91|57\n91|54\n91|74\n91|21\n91|41\n91|62\n91|15\n91|26\n91|94\n91|61\n91|19\n91|25\n91|32\n91|79\n91|42\n91|65\n91|17\n91|13\n91|39\n91|64\n74|54\n74|16\n74|64\n74|57\n74|13\n74|26\n74|15\n74|42\n74|94\n74|61\n74|39\n74|17\n74|32\n74|19\n74|41\n74|79\n74|21\n74|58\n74|73\n74|28\n74|98\n74|25\n74|65\n74|14\n28|11\n28|84\n28|98\n28|37\n28|46\n28|76\n28|38\n28|44\n28|88\n28|24\n28|29\n28|33\n28|67\n28|95\n28|56\n28|55\n28|35\n28|27\n28|97\n28|23\n28|14\n28|34\n28|75\n28|87\n13|24\n13|38\n13|46\n13|23\n13|44\n13|98\n13|87\n13|75\n13|27\n13|84\n13|88\n13|67\n13|95\n13|97\n13|11\n13|28\n13|76\n13|56\n13|29\n13|34\n13|37\n13|14\n13|55\n13|16\n37|39\n37|74\n37|15\n37|41\n37|21\n37|58\n37|33\n37|73\n37|26\n37|64\n37|54\n37|25\n37|65\n37|19\n37|57\n37|61\n37|42\n37|17\n37|32\n37|38\n37|91\n37|46\n37|35\n37|62\n24|46\n24|33\n24|57\n24|62\n24|56\n24|54\n24|84\n24|29\n24|44\n24|97\n24|76\n24|88\n24|87\n24|11\n24|38\n24|95\n24|27\n24|37\n24|34\n24|55\n24|74\n24|75\n24|35\n24|91\n97|15\n97|35\n97|54\n97|88\n97|57\n97|64\n97|19\n97|37\n97|41\n97|21\n97|25\n97|62\n97|91\n97|38\n97|73\n97|32\n97|74\n97|46\n97|42\n97|33\n97|65\n97|61\n97|11\n97|17\n25|13\n25|67\n25|58\n25|19\n25|28\n25|24\n25|75\n25|95\n25|39\n25|16\n25|44\n25|26\n25|41\n25|84\n25|55\n25|79\n25|94\n25|87\n25|98\n25|76\n25|29\n25|23\n25|21\n25|14\n79|75\n79|56\n79|94\n79|23\n79|13\n79|16\n79|76\n79|88\n79|98\n79|27\n79|87\n79|14\n79|24\n79|84\n79|67\n79|28\n79|97\n79|29\n79|37\n79|44\n79|34\n79|55\n79|11\n79|95\n27|15\n27|41\n27|17\n27|11\n27|19\n27|88\n27|64\n27|32\n27|62\n27|65\n27|46\n27|91\n27|38\n27|42\n27|25\n27|35\n27|97\n27|54\n27|73\n27|57\n27|37\n27|74\n27|33\n27|61\n61|16\n61|28\n61|65\n61|24\n61|23\n61|25\n61|19\n61|79\n61|42\n61|39\n61|13\n61|67\n61|32\n61|84\n61|21\n61|41\n61|29\n61|17\n61|94\n61|14\n61|98\n61|55\n61|58\n61|26\n87|56\n87|38\n87|61\n87|73\n87|33\n87|44\n87|54\n87|35\n87|97\n87|88\n87|62\n87|57\n87|95\n87|74\n87|11\n87|75\n87|34\n87|37\n87|76\n87|15\n87|91\n87|64\n87|27\n87|46\n34|27\n34|11\n34|97\n34|91\n34|54\n34|74\n34|42\n34|33\n34|88\n34|56\n34|57\n34|64\n34|17\n34|35\n34|32\n34|61\n34|46\n34|25\n34|38\n34|65\n34|73\n34|37\n34|15\n34|62\n14|35\n14|24\n14|27\n14|56\n14|75\n14|29\n14|87\n14|55\n14|95\n14|76\n14|84\n14|98\n14|97\n14|33\n14|88\n14|91\n14|38\n14|44\n14|46\n14|23\n14|37\n14|11\n14|34\n14|67\n21|44\n21|16\n21|14\n21|67\n21|98\n21|84\n21|27\n21|94\n21|26\n21|76\n21|79\n21|55\n21|75\n21|13\n21|56\n21|28\n21|24\n21|58\n21|39\n21|95\n21|34\n21|29\n21|87\n21|23\n75|88\n75|95\n75|64\n75|35\n75|54\n75|56\n75|27\n75|57\n75|15\n75|46\n75|61\n75|74\n75|33\n75|76\n75|11\n75|73\n75|44\n75|97\n75|42\n75|37\n75|91\n75|38\n75|62\n75|34\n23|55\n23|76\n23|37\n23|44\n23|11\n23|74\n23|84\n23|29\n23|75\n23|33\n23|97\n23|57\n23|95\n23|88\n23|38\n23|27\n23|91\n23|56\n23|46\n23|34\n23|24\n23|62\n23|87\n23|35\n64|32\n64|98\n64|41\n64|25\n64|28\n64|21\n64|61\n64|17\n64|39\n64|24\n64|29\n64|65\n64|19\n64|42\n64|16\n64|13\n64|23\n64|79\n64|58\n64|26\n64|84\n64|94\n64|67\n64|14\n41|16\n41|98\n41|13\n41|67\n41|58\n41|29\n41|87\n41|23\n41|94\n41|28\n41|26\n41|75\n41|84\n41|21\n41|95\n41|44\n41|79\n41|14\n41|39\n41|24\n41|76\n41|34\n41|55\n41|56\n32|41\n32|75\n32|87\n32|29\n32|13\n32|28\n32|24\n32|26\n32|84\n32|17\n32|94\n32|67\n32|21\n32|14\n32|19\n32|98\n32|58\n32|55\n32|79\n32|16\n32|25\n32|39\n32|23\n32|76\n55|56\n55|97\n55|46\n55|27\n55|37\n55|64\n55|11\n55|44\n55|54\n55|57\n55|15\n55|87\n55|74\n55|73\n55|91\n55|75\n55|62\n55|38\n55|33\n55|35\n55|76\n55|34\n55|95\n55|88\n46|79\n46|38\n46|62\n46|35\n46|65\n46|19\n46|91\n46|74\n46|32\n46|15\n46|54\n46|42\n46|73\n46|57\n46|17\n46|25\n46|41\n46|64\n46|58\n46|61\n46|26\n46|39\n46|21\n46|33\n73|98\n73|19\n73|61\n73|79\n73|41\n73|39\n73|64\n73|24\n73|14\n73|16\n73|65\n73|32\n73|42\n73|17\n73|94\n73|25\n73|28\n73|67\n73|13\n73|58\n73|26\n73|29\n73|23\n73|21\n67|37\n67|84\n67|11\n67|97\n67|24\n67|74\n67|55\n67|34\n67|46\n67|91\n67|44\n67|27\n67|29\n67|88\n67|75\n67|35\n67|62\n67|76\n67|33\n67|38\n67|23\n67|95\n67|56\n67|87\n56|33\n56|65\n56|37\n56|35\n56|62\n56|73\n56|91\n56|19\n56|11\n56|25\n56|15\n56|32\n56|38\n56|74\n56|46\n56|27\n56|17\n56|88\n56|64\n56|61\n56|97\n56|42\n56|57\n94|14\n94|87\n94|88\n94|95\n94|23\n94|97\n94|44\n94|28\n94|37\n94|56\n94|84\n94|46\n94|27\n94|13\n94|76\n94|34\n94|98\n94|16\n94|67\n94|24\n94|29\n94|75\n42|29\n42|13\n42|58\n42|65\n42|19\n42|23\n42|98\n42|26\n42|39\n42|87\n42|84\n42|24\n42|25\n42|79\n42|14\n42|16\n42|28\n42|17\n42|55\n42|67\n42|94\n54|23\n54|21\n54|39\n54|26\n54|14\n54|64\n54|28\n54|32\n54|17\n54|25\n54|41\n54|16\n54|15\n54|19\n54|65\n54|67\n54|94\n54|79\n54|58\n54|73\n26|98\n26|27\n26|75\n26|67\n26|94\n26|84\n26|13\n26|87\n26|28\n26|23\n26|16\n26|76\n26|34\n26|56\n26|29\n26|55\n26|88\n26|97\n26|44\n38|15\n38|57\n38|33\n38|25\n38|26\n38|58\n38|64\n38|79\n38|17\n38|35\n38|19\n38|73\n38|42\n38|61\n38|21\n38|62\n38|91\n38|65\n16|95\n16|34\n16|46\n16|27\n16|67\n16|88\n16|87\n16|56\n16|23\n16|11\n16|98\n16|84\n16|76\n16|33\n16|97\n16|28\n16|75\n95|27\n95|15\n95|97\n95|17\n95|73\n95|56\n95|46\n95|42\n95|74\n95|35\n95|38\n95|33\n95|54\n95|61\n95|64\n95|37\n57|26\n57|73\n57|28\n57|67\n57|41\n57|54\n57|58\n57|32\n57|98\n57|79\n57|94\n57|25\n57|13\n57|15\n57|16\n76|37\n76|34\n76|56\n76|74\n76|38\n76|35\n76|65\n76|33\n76|27\n76|46\n76|62\n76|73\n76|88\n76|42\n35|54\n35|61\n35|74\n35|64\n35|73\n35|42\n35|19\n35|15\n35|13\n35|26\n35|65\n35|41\n35|25\n98|24\n98|37\n98|87\n98|27\n98|84\n98|46\n98|34\n98|29\n98|44\n98|91\n98|75\n98|11\n62|39\n62|14\n62|42\n62|57\n62|41\n62|28\n62|94\n62|19\n62|25\n62|64\n62|13\n84|38\n84|97\n84|55\n84|87\n84|56\n84|46\n84|54\n84|57\n84|95\n84|11\n65|17\n65|75\n65|32\n65|84\n65|28\n65|94\n65|55\n65|16\n65|25\n88|33\n88|41\n88|91\n88|32\n88|58\n88|64\n88|35\n88|15\n29|95\n29|55\n29|54\n29|87\n29|74\n29|62\n29|46\n15|94\n15|73\n15|21\n15|79\n15|25\n15|98\n19|84\n19|67\n19|28\n19|55\n19|26\n58|26\n58|44\n58|16\n58|29\n17|44\n17|25\n17|98\n33|32\n33|15\n11|37\n\n38,33,35,91,62,74,57,54,15,73,64,61,65,32,25,19,21,39,58,26,79\n46,62,56,61,35,88,74\n24,17,55,67,42,16,58\n16,39,24,25,28,17,98,41,64,21,29,42,13,79,61,19,14,58,67,94,32\n34,56,97,37,38,33,35,91,62,54,15,73,64,42,65,32,17\n37,46,38,33,91,62,74,57,54,15,73,64,61,42,65,32,17,25,19,41,21,39,58\n65,32,17,25,41,21,39,58,26,79,94,13,16,28,14,98,23,29,84,55,87\n17,25,41,21,58,79,94,13,16,14,98,67,29,55,76\n29,21,25,19,98,94,42,28,26,17,24,55,67,79,39,14,32,23,84,58,13,41,65\n32,42,21,39,79,13,41,94,28\n98,67,23,24,29,84,55,87,75,44,95,34,97,11,37,38,91\n33,94,54,32,42,62,19\n13,95,28,44,24,76,79,39,84,29,98,87,19\n41,21,39,58,26,79,94,13,16,28,14,98,67,23,24,29,84,55,75,76,44,95,34\n58,61,21,13,84,42,79,26,65,14,94,41,23,29,39\n57,15,32,25,21,65,39,79,58,19,73,91,54,26,64,17,61,42,35,33,74\n91,29,75,46,55,11,76,35,84,97,24,95,74,57,33,38,62\n73,55,91,11,33,87,54\n55,24,14,16,23,94,44,98,34,29,76,87,28,84,79,97,95,11,27,26,56\n67,24,55,95,56,27,97,11,88,37,46,38,35,91,62\n16,28,75,21,24,95,84,29,19,44,67,76,87,98,14\n17,25,19,41,21,39,58,79,94,13,16,14,98,67,23,24,84,55,87,75,76\n21,39,26,79,94,13,16,14,98,67,23,24,84,75,76,44,95,34,56\n97,11,88,46,38,33,35,91,62,74,57,54,15,73,64,61,42,65,32,17,25,19,41\n75,95,27,11,88,37,46,35,91,62,74,57,54,15,73,64,61\n32,17,25,19,41,21,39,58,26,79,94,13,16,14,98,67,23,24,29,84,55,87,75\n91,57,15,61,42,17,25,41,21\n95,56,97,11,37,46,38,33,35,91,62,74,57,54,15,73,42,65,32\n38,35,54,88,15,19,39,46,64,65,21,57,73,17,91\n29,84,76,97,37,91,54\n98,67,29,84,55,87,75,76,44,95,34,56,27,97,88,37,46,38,33,35,91\n54,15,73,64,65,32,17,25,19,41,21,39,58,26,16,98,67\n94,13,16,28,98,67,23,29,84,55,87,75,76,44,95,34,56,27,37\n13,16,28,14,98,67,23,24,84,55,87,76,44,56,27,97,11,37,46\n32,21,94,13,28,24,87\n16,14,76,44,97,46,38\n88,25,62,19,65,15,32,57,46,54,39,38,91,42,37,41,35\n95,56,11,88,46,33,91,62,74,57,54,15,73,64,61,42,32\n75,95,56,27,97,37,46,33,35,91,62,57,54,73,61\n29,84,55,87,44,95,34,27,11,88,37,46,38,62,54\n27,11,91,61,65,32,17\n76,44,95,34,56,27,97,88,37,46,38,33,91,62,74,57,54,15,73,61,42\n54,73,64,61,65,32,25,19,41,21,39,58,94,13,16,98,67\n28,23,84,76,44,88,37,46,33\n14,84,55,87,44,88,35\n57,11,56,84,75,35,33,54,15,27,95,34,55\n61,46,38,91,73,11,17,64,65,74,35,57,37,34,32\n14,23,24,29,84,55,75,76,95,34,56,27,97,88,37,46,38,33,35\n91,33,15,62,64,41,17,32,54,42,25,65,35,61,88,38,74,37,57,97,73\n27,97,33,62,74,54,64,65,19\n44,34,88,37,54,64,65\n16,56,44,26,14,76,34,94,28,75,27,95,67,98,23,79,87,97,58\n42,21,32,94,41,79,26,25,57,65,33,62,54,15,91,74,35\n62,54,19,39,16\n97,61,44,54,62,95,56,46,27,37,34,57,88,38,91\n61,65,17,25,19,94,13,28,84\n35,64,32,42,37,65,25,57,33,73,17,11,46,19,54,62,88,41,38\n73,25,61,42,17,41,62,35,88,65,54\n87,76,44,95,56,11,88,37,46,33,91,62,15,73,64\n54,15,39,62,21,32,61,46,41,74,33\n76,95,34,88,37,46,33,35,62,57,54,15,64,61,42\n25,57,65,37,17,56,15\n21,58,26,16,14,67,29,84,76,44,56\n74,57,54,15,73,64,61,42,65,32,17,25,19,41,39,58,26,79,94,13,16,28,14\n23,24,84,55,76,95,27,97,11,88,37,46,38,35,91,62,74\n26,73,21,25,35,54,74,17,64,94,62\n16,61,19,23,65,32,25,39,15\n32,25,19,41,21,58,26,94,13,16,14,98,67,23,24,29,84,55,75\n34,55,95,75,57,56,97,38,11,84,29,76,91,37,74,87,62,24,27,33,35\n21,55,25,67,17,16,29,19,98,58,24,32,26,87,84,75,23,41,14\n61,28,94,14,25,79,65,98,41,67,26,21,13,24,19,32,58,29,39,84,42\n42,32,17,25,19,41,21,58,94,13,28,14,98,67,23,24,29\n32,26,87,21,39,41,79,55,98,67,24,17,16,75,28\n84,21,28,67,87,29,76,79,95,41,14,55,39,23,26,44,13,16,94,98,24\n39,79,94,13,28,98,67,23,24,29,84,55,87,75,76,95,34,56,27\n34,15,57,76,91,75,37,38,54,55,27,44,33,46,88,87,62,73,11\n55,87,75,76,95,34,27,97,11,88,38,35,91,74,57,54,73\n97,37,38,62,74,57,15,73,64,42,17,25,41\n95,34,56,97,11,88,37,46,33,35,91,74,57,15,73,64,61,65,32\n21,58,26,13,16,28,14,98,67,23,24,29,84,87,76,44,95,34,56\n97,98,11,38,87,23,16\n26,19,84,76,58,79,94,25,23,28,98,87,41,39,75,13,44\n95,34,56,27,97,11,88,37,46,38,91,62,74,57,15,64,61,42,65\n75,44,95,34,56,11,88,37,35,73,64\n15,73,64,61,42,32,17,25,19,21,39,58,26,79,94,13,16,28,14,98,23\n79,13,16,14,98,23,24,29,84,55,87,75,76,95,34,56,11\n94,13,16,28,98,24,29,84,55,87,75,76,95,34,27,97,11\n94,98,67,76,56,97,37\n19,21,39,13,16,23,24,29,84,55,75,76,95\n54,65,32,17,25,26,79,13,16,14,67\n27,91,74,57,65,17,19\n38,33,35,91,62,54,15,64,61,42,65,32,17,25,19,41,21,39,79\n16,28,23,24,55,87,44,34,56,27,97,11,37,46,38\n54,15,44,74,65,38,95,62,73,64,42,33,34\n16,14,98,67,24,29,55,87,75,76,44,97,11,88,37,46,38\n73,57,42,54,27,33,11,17,37,64,65,62,97\n91,57,54,15,73,64,42,65,17,25,19,21,39,26,79,94,16\n41,39,58,26,79,94,13,16,28,14,98,67,23,24,29,84,55,87,75,76,44,95,34\n28,14,98,67,23,24,29,84,55,87,75,76,44,95,34,56,27,11,88,37,46,38,33\n46,33,74,88,64,62,27,37,65,38,44\n61,42,65,32,17,25,19,41,39,58,26,79,94,13,16,28,98,67,24,29,84\n64,14,39,32,61,28,57,13,25,16,17,65,94,58,21,41,19,79,98,15,73,26,54\n91,46,65,26,19,74,41\n23,24,29,84,55,87,75,76,44,95,34,56,27,11,37,46,38,33,35,62,74\n64,61,42,65,32,17,25,19,41,21,39,58,26,79,94,13,16,28,14,98,67,24,29\n19,21,13,16,23,75,95\n37,46,15,73,64,61,32\n87,97,46,38,57\n38,57,46,75,29,91,56,97,11,55,84,27,76,37,54,87,44,74,95\n73,61,42,65,32,39,24\n24,29,84,55,75,95,34,56,27,97,11,88,37,46,33,91,62,74,57\n79,94,28,98,67,23,24,84,55,87,76,95,34,56,88\n33,27,62,64,61,88,95,97,54,74,56\n98,46,97,16,88,95,56,14,13\n15,17,61,73,88,41,19,25,37,62,64,35,54,65,33,11,46,57,97\n13,14,23,87,95,27,11\n55,13,25,79,29,14,42,94,84,39,24,67,19\n88,29,34,67,28,23,79,87,24\n41,21,39,58,26,79,94,13,16,14,98,67,23,24,29,84,55,87,75,76,44,95,34\n16,65,17,19,25\n64,97,25,57,35,54,74,19,46,65,17,32,11,91,61,15,62,73,41,42,33,88,38\n65,32,17,25,39,58,26,14,98,23,29,84,87\n57,55,95,56,76,62,75,15,74,27,88,73,54,87,37,46,38\n17,11,46,25,21,91,73\n67,87,75,76,34,46,35,91,62\n57,37,33,65,42,61,58,39,17,41,64,21,15,32,46,25,19,35,73,91,62\n17,28,39,58,19,32,24,41,13,14,23,26,29,98,16,84,42\n21,67,23,41,64,98,13,17,16,24,32,73,14,26,19\n55,16,95,21,56,26,94,79,58\n76,44,33,91,74\n55,87,75,44,34,27,88,37,33,91,62,74,54,15,73\n39,41,54,88,19,62,61\n42,64,21,79,17,25,98,57,39,41,28,54,61,94,15,32,19,58,14,13,73\n65,95,61,42,73,34,38,56,64,54,44\n25,64,28,39,16,62,32\n21,39,58,79,13,16,24,84,95,34,56\n34,97,88,91,62,74,54,15,73,61,17\n73,46,54,88,74,61,15,33,11,62,65,37,41,42,17,38,21,19,57,32,25,35,64\n91,54,15,73,64,61,42,65,17,19,21,39,79\n75,76,56,27,97,11,37,38,35,91,57,15,61\n41,21,39,26,94,13,16,28,14,98,67,23,29,84,55,87,75,76,44,95,34\n28,14,67,23,29,84,55,87,75,76,95,34,27,97,11,88,46,38,33\n95,97,46,91,62,74,57,73,61,65,32\n94,13,16,28,14,98,67,23,24,29,55,87,75,76,44,95,34,56,27,97,11,88,37\n67,24,29,84,95,56,11,88,37,46,33,35,91\n37,46,38,33,62,74,57,15,73,61,42,19,58\n19,21,39,58,79,13,98,76,95\n27,79,94,24,87,75,88,55,14\n91,88,29,34,98,75,56,23,97,33,95,84,27,46,67,44,87,38,76,37,35\n79,16,28,98,23,84,75,97,88\n23,13,39,28,76,16,56,26,98,44,34,87,84,29,58,79,27\n37,46,38,33,35,91,74,57,54,15,64,42,65,32,17,25,19,41,21,39,58\n33,97,76,14,29,98,28,34,56,27,87,38,88,24,23,11,46\n28,23,55,87,76,34,56,11,88,37,46,38,33\n62,74,73,64,61,42,65,41,21,39,79,94,28\n62,74,57,54,15,64,61,42,65,32,17,25,21,39,58,26,79,94,13,16,28\n26,79,16,28,14,24,84,87,76,95,11\n67,23,24,84,55,87,75,44,34,46,33,91,62\n41,21,39,58,26,79,94,13,28,14,98,67,23,24,29,84,87,75,76,44,34\n16,67,24,29,55,34,11\n27,98,37,91,55,44,95\n91,62,74,57,15,73,64,61,42,17,19,41,39,26,94,13,16\n26,98,29,84,55,95,27\n55,87,76,44,95,34,56,27,97,88,38,33,35,91,62\n54,91,42,79,57,62,35\n26,79,13,14,98,67,24,84,55,75,44,95,34,97,11\n28,98,19,26,39,75,84,32,67,29,55,24,17,41,23\n33,35,91,62,57,54,15,73,64,61,42,65,32,17,25,19,41,21,39,58,26,79,94\n15,98,42,65,64,17,94,28,58,39,57,32,25,41,73,19,16\n16,98,67,23,29,84,55,87,75,44,95,34,56,27,11,88,37,46,38\n17,73,56,42,11,34,37,46,61,97,54\n29,84,75,56,37,62,54\n37,46,38,33,35,62,74,57,54,15,73,64,61,42,65,32,17,25,19,41,21,39,58\n91,21,13,41,64,57,15,42,26,79,19,65,32,54,35,25,94,39,62,73,58,17,74\n62,74,54,15,42,65,17,25,19,41,39,16,28\n91,62,74,57,64,42,65,32,19,21,39,26,79,94,16\n75,88,46,76,11,38,67,95,33,91,97,23,56,44,34,62,87,84,29\n73,16,79,67,26,65,21,94,64,28,61,39,58,19,15\n33,35,91,62,74,57,73,65,25,19,41,21,94\n75,76,44,95,34,56,27,97,11,88,46,38,35,91,62,54,73,64,61\n41,21,94,13,14,67,23,84,55,87,75,44,34\n65,25,28,98,87\n16,28,14,98,29,55,87,75,97,37,38\n55,76,95,97,88,37,38,35,62,57,15\n17,25,19,41,21,39,58,26,79,94,13,16,28,14,98,23,24,29,84,55,87,75,76\n65,25,41,21,39,79,94,14,67,24,84,55,87\n61,19,13,14,29\n24,29,84,87,75,95,34,11,37,46,35,91,62,74,57\n34,56,27,97,88,37,46,33,91,62,74,57,54,15,73,64,61,42,65,32,17\n15,64,42,32,17,25,41,21,39,58,26,79,94,13,16,28,67\n21,58,16,14,84,87,75\n46,33,35,62,54,73,65,17,25,39,26\n56,97,37,46,38,35,91,62,57,54,15,61,65,32,17\n19,15,79,98,64,21,16,67,42,73,17,13,26,28,41"

	dependents, updates := parseInput(input)
	total := calculateTotal(updates, dependents)

	fmt.Printf("Total: %d\n", total)
}

func calculateTotal(updates [][]int, dependents map[int][]int) int {
	total := 0

	for _, update := range updates {
		//fmt.Printf("Checking update %d (%#v)\n", i+1, update)
		total += getValue(update, dependents, true)
	}

	return total
}

func getValue(update []int, dependents map[int][]int, firstTry bool) int {
	seenValues := map[int]int{}

	for i, value := range update {
		for _, dependent := range dependents[value] {
			if seenIndex, seen := seenValues[dependent]; seen {
				//fmt.Printf("Not correct: %d must come before %d\n", value, dependent)
				var fixedUpdate []int
				if seenIndex > 0 {
					fixedUpdate = append(fixedUpdate, update[0:seenIndex]...)
				}
				fixedUpdate = append(fixedUpdate, value)
				fixedUpdate = append(fixedUpdate, update[seenIndex:i]...)
				fixedUpdate = append(fixedUpdate, update[i+1:]...)
				//fmt.Printf("Fixed %#v to %#v\n", update, fixedUpdate)
				return getValue(fixedUpdate, dependents, false)
			}
		}
		//fmt.Printf("Seen %d\n", value)
		seenValues[value] = i
	}

	//fmt.Println("Valid!")
	if firstTry {
		return 0
	} else {
		return update[len(update)/2]
	}
}

func parseInput(input string) (map[int][]int, [][]int) {
	dependents := map[int][]int{}
	var updates [][]int

	split := strings.Split(input, "\n\n")

	for _, line := range strings.Split(split[0], "\n") {
		var first, second int
		_, _ = fmt.Sscanf(line, "%d|%d", &first, &second)

		_, exists := dependents[first]
		if exists {
			dependents[first] = append(dependents[first], second)
		} else {
			dependents[first] = []int{second}
		}
	}

	for _, line := range strings.Split(split[1], "\n") {
		var update []int
		for _, val := range strings.Split(line, ",") {
			intVal, _ := strconv.Atoi(val)
			update = append(update, intVal)
		}
		updates = append(updates, update)
	}

	return dependents, updates
}
