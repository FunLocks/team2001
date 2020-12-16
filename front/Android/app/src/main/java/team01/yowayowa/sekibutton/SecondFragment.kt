package team01.yowayowa.sekibutton

import android.content.Intent
import android.net.Uri
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Button
import android.widget.Toast
import androidx.constraintlayout.widget.ConstraintLayout
import androidx.navigation.fragment.findNavController

/**
 * A simple [Fragment] subclass as the second destination in the navigation.
 */
class SecondFragment : Fragment() {

    override fun onCreateView(
            inflater: LayoutInflater, container: ViewGroup?,
            savedInstanceState: Bundle?
    ): View? {
        // Inflate the layout for this fragment
        return inflater.inflate(R.layout.fragment_second, container, false)
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        val webButton = view.findViewById<ConstraintLayout>(R.id.webButton)
        webButton.setOnClickListener {
            val url : String = "https://github.com/FunLocks/team2001"
            val uri : Uri = Uri.parse(url)
            val intent :Intent = Intent(Intent.ACTION_VIEW,uri)
            startActivity(intent)
            Toast.makeText(context,"クリックされました",Toast.LENGTH_SHORT).show()
        }
    }
}